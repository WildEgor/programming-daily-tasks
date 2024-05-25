package main

import (
	"fmt"
	"sync"
)

// HINT: https://youtu.be/R68oxsnnJBE?si=0NC_IBf_NV6EwUij&t=692

func main() {
	var mu sync.Mutex
	var max int // HINT: Race-condition, 'cause not safe read/write for many goroutines

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 1000; i > 0; i-- {
		// i := i // HINT: depends on golang version, suitable for go <= 1.20

		go func(i int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			if i%2 == 0 && i > max {
				max = i
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("Maximum is %d", max)
}
