package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	data  chan int
	total chan int
}

func NewCounter() *Counter {
	c := &Counter{make(chan int), make(chan int)}

	go func() {
		var count int
		for {
			select {
			case increment := <-c.data:
				count += increment
			case c.total <- count:
			}
		}
	}()

	return c
}

func (c *Counter) Add(v int) {
	c.data <- v
}

func (c *Counter) Total() int {
	return <-c.total
}

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(9)

	for i := 0; i < 9; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println(counter.Total())
}
