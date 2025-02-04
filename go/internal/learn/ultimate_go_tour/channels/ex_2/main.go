// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

import (
	"fmt"
	"math/rand"
)

// Add imports.

// Declare constant for number of goroutines.
const gsn = 100

func init() {
	// Seed the random number generator.
	rand.New(rand.NewSource(1))
}

func main() {

	// Create the buffered channel with room for
	// each goroutine to be created.
	ch := make(chan int, gsn)

	// Iterate and launch each goroutine.
	for i := 0; i < gsn; i++ {
		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			ch <- rand.Int()
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	wait := gsn

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	store := make([]int, 0)
	for wait > 0 {
		store = append(store, <-ch)
		wait--
	}

	// Print the values in our slice.
	fmt.Println(store)
}
