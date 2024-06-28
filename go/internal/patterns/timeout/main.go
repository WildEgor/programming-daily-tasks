package main

import (
	"fmt"
	"github.com/wildegor/daily_tasks/go/internal/patterns/timeout/utils_timeout"
	"time"
)

func doJob() error {
	// Simulate work
	time.Sleep(2999 * time.Millisecond)
	fmt.Println("OK!")
	return nil
}

func main() {
	// Change timeout to 4 seconds to prevent timeout
	if err := utils_timeout.WithTimeout(3*time.Second, doJob); err != nil {
		fmt.Println(err)
	}
}
