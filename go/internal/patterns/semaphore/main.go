package main

import (
	"fmt"
	"sync"
	"time"
)

type Sem struct {
	ch chan struct{}
}

func NewSem(n int) *Sem {
	return &Sem{
		ch: make(chan struct{}, n),
	}
}

func (s *Sem) Acquire() {
	s.ch <- struct{}{}
}

func (s *Sem) Release() {
	<-s.ch
}

func main() {
	const numWorkers = 20
	const numTokens = 5

	sem := NewSem(numTokens)

	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	callback := func(id int) {
		sem.Acquire()
		defer sem.Release()
		defer wg.Done()

		time.Sleep(time.Second)
		fmt.Printf("Goroutine %d is acquired the token.\n", id)
	}

	for i := 1; i <= numWorkers; i++ {
		go callback(i)
	}

	wg.Wait()
}
