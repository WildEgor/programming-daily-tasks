package main

import (
	"fmt"
	"sync"
)

func joinChannels(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	go func() {
		for _, ch := range chs {
			go func(ch <-chan int) {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}(ch)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func spawn() <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range []int{1, 0, 1} {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func main() {
	chA := spawn()
	chB := spawn()
	chC := spawn()

	chOut := joinChannels(chA, chB, chC)

	r := make([]int, 0)
	for v := range chOut {
		r = append(r, v)
	}

	fmt.Println(r)
}
