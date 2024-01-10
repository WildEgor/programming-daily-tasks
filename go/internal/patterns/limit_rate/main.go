package main

import (
	"fmt"
	"github.com/wildegor/daily_tasks/go/internal/patterns/limit_rate/limiter"
	"net/http"
	"time"
)

func someHandler() bool {
	time.Sleep(10 * time.Second)
	return true
}

func WithLimit(handler http.HandlerFunc, l *limiter.Limiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := l.Process(func() {
			handler(w, r)
		})

		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many requests"))
		}
	}
}

func main() {
	l := limiter.NewLimiter(10)

	http.HandleFunc("/", WithLimit(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := someHandler()
		w.Write([]byte(fmt.Sprintf("result: %v", result)))
	}), l))

	http.ListenAndServe(":8080", nil)
}
