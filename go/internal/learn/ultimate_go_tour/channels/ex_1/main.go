// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Add imports.

func main() {

	// Create an unbuffered channel.
	ch := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		defer wg.Done()
		goroutine("1", ch)
	}()

	// Launch the goroutine and handle Done.
	go func() {
		defer wg.Done()
		goroutine("2", ch)
	}()

	// NotifyUser a value to start the counting.
	ch <- 0

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(name string, ch chan int) {
	now := time.Now().Unix()

	for {

		// Wait for the value to be sent.
		// If the channel was closed, return.
		v, ok := <-ch
		if !ok {
			fmt.Printf("Goroutine(%s) Down at %v!\n", name, now)
			return
		}

		// Display the value.
		fmt.Printf("Goroutine(%s) Inc %d at %v\n", name, v, now)

		// Terminate when the value is 10.
		if v == 10 {
			close(ch)
			fmt.Printf("Goroutine(%s) with ch closed at %v!\n", name, now)
			return
		}

		// Increment the value and send it
		// over the channel.
		v++
		ch <- v
	}
}
