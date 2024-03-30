// Fix the race condition in this program.
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// numbers maintains a set of random numbers.
var numbers []int

func main() {

	// Number of goroutines to use.
	const grs = 3

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	var m sync.Mutex

	// Create three goroutines to generate random numbers.
	for i := 0; i < grs; i++ {
		go func() {
			random(&m, 10)
			wg.Done()
		}()
	}

	// Wait for all the goroutines to finish.
	wg.Wait()

	// Display the set of random numbers.
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}

// random generates random numbers and stores them into a slice.
func random(m *sync.Mutex, amount int) {
	m.Lock()
	defer m.Unlock()

	// Generate as many random numbers as specified.
	for i := 0; i < amount; i++ {
		n := rand.Intn(100)
		numbers = append(numbers, n)
	}
}
