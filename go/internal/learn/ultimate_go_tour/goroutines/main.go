// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

// Add imports.
import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	// Declare an anonymous function and create a goroutine.
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		defer wg.Done()
		for i := 100; i >= 0; i-- {
			fmt.Println(i)
		}
	}()

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display "Terminating Program".
	fmt.Println("Terminating Program")
}
