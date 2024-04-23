package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	req, _ := http.NewRequest("GET", "http://google.com", nil)
	req = req.WithContext(ctx)

	ch := make(chan error)

	go func() {
		_, err := http.DefaultClient.Do(req)

		select {
		case <-ctx.Done():
			fmt.Println("timeout exceed!")
		default:
			ch <- err
		}
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	select {
	case err := <-ch:
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("no error")
	case <-ctx.Done():
		fmt.Println("context canceled!")
	}

}
