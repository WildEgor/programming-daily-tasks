package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// HINT: https://youtu.be/R68oxsnnJBE?si=v6ye3VXcHv33mXtA
// RPS?
// Cache for n day?

// Update weather async

type weatherFetcherFn func() int

type WeatherCache struct {
	mu    sync.RWMutex // HINT: go uber style guide
	value int

	callback weatherFetcherFn
}

func NewWeatherCache(callback weatherFetcherFn) *WeatherCache {
	wc := &WeatherCache{
		callback: callback,
	}

	wc.update()

	return wc
}

func (wc *WeatherCache) Set(data int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	wc.value = data
}

func (wc *WeatherCache) Get() int {
	wc.mu.RLock()
	defer wc.mu.RUnlock()
	return wc.value
}

func (wc *WeatherCache) update() {
	go func() {
		for range time.Tick(1 * time.Second) {
			if wc.callback != nil {
				fmt.Println("fetch")

				v := wc.callback()
				wc.Set(v)
			}
		}
	}()
}

func getWeather() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func main() {
	mux := http.NewServeMux()

	cache := NewWeatherCache(getWeather)

	mux.HandleFunc("/weather/highload", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "{\"result\":%d}\n", cache.Get())
	})

	srv := &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
