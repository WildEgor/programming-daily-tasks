package main

import (
	"fmt"
	"sync"
)

func joinChannels(chs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	wg := &sync.WaitGroup{}
	wg.Add(len(chs))

	go func() {
		for _, ch := range chs {
			ch := ch // HINT: if using Go < 1.22
			go func() {
				defer wg.Done()

				for v := range ch {
					out <- v
				}
			}() // HINT: or can provide using func args
		}

		wg.Wait()

		close(out)
	}()

	return out
}

func spawn(nums []int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func main() {
	ch1 := spawn([]int{1, 2, 3})
	ch2 := spawn([]int{4, 5, 6})
	ch3 := spawn([]int{7, 8, 9})

	for num := range joinChannels(ch1, ch2, ch3) {
		fmt.Printf("%d\n", num)
	}
}
