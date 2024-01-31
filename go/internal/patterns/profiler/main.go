package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Регистрируем обработчики профайлера
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	http.ListenAndServe(":8080", nil)
}
