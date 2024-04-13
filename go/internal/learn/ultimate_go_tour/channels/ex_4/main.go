// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Add imports.

func main() {

	// Create the channel for sharing results.
	share := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	done := make(chan bool)

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.
	ps := runtime.GOMAXPROCS(0)
	now := time.Now()

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	wg := new(sync.WaitGroup)
	wg.Add(ps)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := 0; i < ps; i++ {
		go func(id int) {

			// Start an infinite loop.
			for {

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select {
				// In one case send the random number.
				case share <- n:
					fmt.Printf("Worker %d sent %d\n", id, n)
				// In another case receive from the shutdown channel.
				case <-done:
					fmt.Printf("Worker %d shutting down\n", id)
					wg.Done()
					return
				}
			}
		}(i)
	}

	// Create a slice to hold the random numbers.
	nums := make([]int, 0)

	// Receive from the values channel with range.
	for v := range share {
		// continue the loop if the value was even.
		if v%2 == 0 {
			continue
		}
		// Store the odd number.
		nums = append(nums, v)
		// break the loop once we have 100 results.
		if len(nums) == 100 {
			break
		}
	}

	// NotifyUser the shutdown signal by closing the shutdown channel.
	close(done)

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Print the values in our slice.
	took := time.Since(now)
	fmt.Printf("Took %s\n", took)
	fmt.Printf("GOMAXPROCS use %d\n", ps)
	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}
