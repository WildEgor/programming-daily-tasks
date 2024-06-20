package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	cancelFunc context.CancelFunc
)

func killThemAll() {
	if cancelFunc != nil {
		fmt.Println("FINISH HIM!")
		cancelFunc()
	}
}

func worker(ctx context.Context, indx int, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()

	time.Sleep(3 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d killed\n", indx)
		return
	case <-ch:
		fmt.Println(indx)
	}
}

func main() {
	const n = 10
	const p = 5

	wg := &sync.WaitGroup{}
	ch := make(chan struct{}, p) // act like semaphore

	ctx, cancel := context.WithCancel(context.Background())
	cancelFunc = cancel

	for i := 1; i <= n; i++ {
		wg.Add(1)
		ch <- struct{}{}

		go worker(ctx, i, wg, ch)
	}

	time.Sleep(1 * time.Second)
	killThemAll()

	wg.Wait()

	fmt.Println("All workers finished")
}
