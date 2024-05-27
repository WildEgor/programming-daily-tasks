package main

import (
	"fmt"
	"math"
	"sync"
)

func joinChannels(chs ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		var wg sync.WaitGroup

		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()

				for v := range ch {
					out <- v
				}
			}(ch, &wg)
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

func binaryToDecimal(binary []int) int {
	decimal := 0
	length := len(binary)

	for i, bit := range binary {
		if bit == 1 {
			// The power of 2 for the current bit position is calculated as (length-1)-i
			decimal += int(math.Pow(2, float64((length-1)-i)))
		}
	}
	return decimal
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

	fmt.Println(binaryToDecimal(r))
}
