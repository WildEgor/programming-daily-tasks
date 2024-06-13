package main

import (
	"fmt"
	"sync"
)

func fanIn(inps ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(inps))
	for _, c := range inps {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func main() {
	c1 := generator(1, 3, 5)
	c2 := generator(2, 4, 6)

	c := fanIn(c1, c2)

	for n := range c {
		fmt.Println(n)
	}
}
