package main

import (
	"context"
	"fmt"
	"time"
)

func Process(ctx context.Context) {
	for {
		time.Sleep(time.Second * 10)
		// time.Sleep(10 * time.Millisecond)

		select {
		case <-ctx.Done():
			fmt.Println("stop goroutine")
			break
		}
	}
}

func main() {
	startTime := time.Now()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	go Process(ctx)

	<-ctx.Done()

	duration := time.Since(startTime)
	fmt.Printf("Duration: %s\n", duration)

	fmt.Print("context done")
}
