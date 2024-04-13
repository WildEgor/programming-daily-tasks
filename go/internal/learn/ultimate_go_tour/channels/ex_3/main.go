// Write a program that uses goroutines to generate up to 100 random numbers.
// Do not send values that are divisible by 2. Have the main goroutine receive
// values and add them to a slice.
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Declare constant for number of goroutines.
const goroutines = 100

func init() {
	// Seed the random number generator.
	rand.New(rand.NewSource(1))
}

func main() {

	// Create the channel for sharing results.
	share := make(chan int)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	wg := new(sync.WaitGroup)
	wg.Add(goroutines)

	// Iterate and launch each goroutine.
	for i := 0; i < goroutines; i++ {
		// Create an anonymous function for each goroutine.
		go func(g *sync.WaitGroup) {
			// Ensure the waitgroup is decremented when this function returns.
			defer g.Done()

			// Generate a random number up to 1000.
			n := rand.Intn(1000)

			// Return early if the number is even. (n%2 == 0)
			if n%2 == 0 {
				return
			}

			// NotifyUser the odd values through the channel.
			fmt.Printf("NotifyUser value %d\n", n)
			share <- n
		}(wg)
	}

	// Create a goroutine that waits for the other goroutines to finish then
	// closes the channel.
	go func() {
		defer close(share)
		fmt.Println("Close ch!")
		wg.Wait()
	}()

	// Receive from the channel until it is closed.
	// Store values in a slice of ints.
	nums := make([]int, 0)
	for v := range share {
		fmt.Printf("Recieve value %d\n", v)
		nums = append(nums, v)
	}

	// Print the values in our slice.
	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}
