package main

import (
	"fmt"
	"github.com/wildegor/daily_tasks/go/internal/patterns/timeout/timeout"
	"time"
)

func someHandler() error {
	time.Sleep(3 * time.Second)
	fmt.Sprintf("result: %v", true)
	return nil
}

func main() {
	err := timeout.WithTimeout(2*time.Second, someHandler) // Change timeout to 4 seconds to prevent timeout
	if err != nil {
		fmt.Println(err)
	}
}
