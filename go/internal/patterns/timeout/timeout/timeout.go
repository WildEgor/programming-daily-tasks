package timeout

import (
	"errors"
	"time"
)

func WithTimeout(dur time.Duration, f func() error) error {
	var err error
	done := make(chan struct{})
	go func() {
		err = f()
		close(done)
	}()

	select {
	case <-done:
		return err
	case <-time.After(dur):
		return errors.New("timeout")
	}
}
