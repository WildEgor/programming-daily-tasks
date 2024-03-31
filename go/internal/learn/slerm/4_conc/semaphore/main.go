package main

import (
	"fmt"
	"time"
)

/*
Напишите программу, которая реализует rate limit на основе семафора.
Например, есть 1000 запросов, которые нужно отправить по сети, но одновременно не должно выполняться более 10 запросов, чтобы не перегрузить сервер.
В качестве выполнения запросов так же используйте time.Sleep.
Выведите в консоль тайминги выполнения каждого запроса и убедитесь, что в каждую секунду выполняется не более 10 горутин.
P.S. Да, семафор уже есть в golang встроенный. Его реализацию можно посмотреть в исходнике в конце страницы https://pkg.go.dev/golang.org/x/sync/semaphore.
Она сложнее той, что предлагается написать вам.
*/

type semType chan struct{}

func NewSemaphore(n int) semType {
	return make(semType, n)
}

func (s semType) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s semType) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

const (
	N      = 10
	Total  = 20
	Timout = 1 * time.Second
)

func process(id int) {
	fmt.Printf("[%s]: running task %d\n", time.Now().Format("15:04:05"), id)
	time.Sleep(Timout)
}

func main() {
	sem := NewSemaphore(N)
	done := make(chan bool)

	for i := 1; i <= Total; i++ {
		go func(v int) {
			sem.Acquire(1)

			go func() {
				defer sem.Release(1)

				process(v)

				// Last task
				if v == Total {
					done <- true
				}
			}()

		}(i)
	}

	<-done
}
