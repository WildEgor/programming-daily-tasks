package main

import (
	"fmt"
	"sync"
	"time"
)

func WithTime(name string, fn func()) {
	startTime := time.Now()
	fn()
	duration := time.Since(startTime)
	fmt.Printf("Duration with \"%s\": %s\n", name, duration)
}

type Valuable interface {
	value() int
}

type Small struct {
	v int
}

func (s Small) value() int {
	return s.v
}

type Big struct {
	v int

	name   string
	email  string
	age    int
	gender string
	cityId string
}

func (b Big) value() int {
	return b.v
}

func measureChannels(ref interface{}) {
	ch := make(chan interface{}, 10000000)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()

		for range 10000000 {
			ch <- ref
		}

		close(ch)
	}()

	wg.Wait()

	var sum int
	for val := range ch {
		if v, ok := val.(Valuable); ok {
			sum += v.value()
		}
	}

	fmt.Println(sum)
}

func main() {
	WithTime("test small", func() {
		measureChannels(Small{v: -1})
	})

	WithTime("test big", func() {
		measureChannels(Big{v: 1})
	})
}
