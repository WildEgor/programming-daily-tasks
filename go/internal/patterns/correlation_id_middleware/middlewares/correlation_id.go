package middlewares

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type correlationIdKey int

const key correlationIdKey = 1
const correlationIdHeader = "x-correlation-id"

func contextWithCorrelationID(ctx context.Context, uuid string) context.Context {
	return context.WithValue(ctx, key, uuid)
}

func ExtractCorrelationIDFromContext(ctx context.Context) (string, bool) {
	cid, ok := ctx.Value(key).(string)
	return cid, ok
}

type RequestDecorator func(r *http.Request) *http.Request

func RequestWithCorrelationId(r *http.Request) *http.Request {
	ctx := r.Context()

	if uid, ok := ExtractCorrelationIDFromContext(ctx); ok {
		r.Header.Add(correlationIdHeader, uid)
	}

	return r
}

func CorrelationIDMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if uid := request.Header.Get(correlationIdHeader); uid != "" {
			ctx = contextWithCorrelationID(ctx, uid)
		} else {
			ctx = contextWithCorrelationID(ctx, uuid.NewString())
		}

		request = request.WithContext(ctx)
		h.ServeHTTP(writer, request)
	})
}

type Logger interface {
	Log(context.Context, string)
}

type MyLogger struct {
}

func (l *MyLogger) Log(ctx context.Context, msg string) {
	if uid, ok := ExtractCorrelationIDFromContext(ctx); ok {
		msg = fmt.Sprintf("LOG (%s) - %s", uid, msg)
	}

	fmt.Println(msg)
}
