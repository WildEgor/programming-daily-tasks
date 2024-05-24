package main

import (
	"fmt"
	"runtime"
	"time"
)

//func main() {
//	ch := make(chan int, 1) // HINT: buffered channel
//	defer close(ch)         // HINT: close channel after done main
//
//	go func(ch chan int) {
//		time.Sleep(time.Second)
//		<-ch
//	}(ch) // HINT: "run" (will be scheduled) goroutine
//
//	ch <- 1 // HINT: not blocking, if not buffered alseep deadlock 'cause second write to channel block no any reader
//	ch <- 2
//}

func main() {
	runtime.GOMAXPROCS(1) // HINT: pure concurrency 'cause goroutines running only at one core

	go func() {
		var u int

		for {
			u -= 2
			fmt.Println(u)

			if u == 1 {
				break
			}
		}
	}()

	<-time.After(time.Second * 5) // HINT: changed to seconds for demo and blocked

	fmt.Println("Hello world from main.go")
}
