package main

import "fmt"

// The generator pattern is a way to create a function that gives us a series of values. In other programming languages,
// people might use things called enumerators or iterators. But in Go, the best tool for this job is a channel.
// For Node.js analogy is Generator (with yield)

// evenGenerator lets us handle the data one piece at a time. It gives us each value only when we ask for it.
// This is especially useful when we're dealing with really big amounts of data that won't fit in our computer's memory all at once.
func evenGenerator(n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		if n <= 0 {
			out <- 0
		}

		for i := 0; i < n; i += 2 {
			out <- i
		}
	}()

	return out
}

func main() {
	for v := range evenGenerator(-1) {
		fmt.Println(v)
	}
}
