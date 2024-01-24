package main

import (
	"fmt"
	"github.com/wildegor/daily_tasks/go/internal/patterns/correlation_id_middleware/middlewares"
	"net/http"
)

func main() {
	logger := &middlewares.MyLogger{}

	http.Handle("/log", middlewares.CorrelationIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Log(ctx, "Incoming request")

		w.Write([]byte(fmt.Sprintf("result: %v", "")))
	})))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("bootstrap error")
	}
}
