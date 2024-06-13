package utils_timeout

import (
	"errors"
	"time"
)

func WithTimeout(dur time.Duration, f func() error) error {
	var err error
	done := make(chan struct{})
	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	// Call callback in goroutine
	go func() {
		err = f()
		close(done)
	}()

	select {
	case <-done:
		return err
	case <-ticker.C:
		return errors.New("timeout")
	}
}
