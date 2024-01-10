package limiter

import "fmt"

type Limiter struct {
	ch chan struct{}
}

func NewLimiter(limit int) *Limiter {
	ch := make(chan struct{}, limit)

	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}

	return &Limiter{ch: ch}
}

func (r *Limiter) Wait() {
	<-r.ch
}

func (r *Limiter) Release() {
	r.ch <- struct{}{}
}

func (r *Limiter) Close() {
	close(r.ch)
}

func (r *Limiter) Closed() bool {
	select {
	case <-r.ch:
		return true
	default:
		return false
	}
}

func (r *Limiter) Limit() int {
	return cap(r.ch)
}

func (r *Limiter) Available() int {
	return len(r.ch)
}

func (r *Limiter) Used() int {
	return r.Limit() - r.Available()
}

func (r *Limiter) String() string {
	return fmt.Sprintf("Limiter{limit: %d, available: %d}", r.Limit(), r.Available())
}

func (r *Limiter) Do(f func()) error {
	r.Wait()
	defer r.Release()
	f()
	return nil
}

func (r *Limiter) Process(f func()) error {
	select {
	case <-r.ch:
		defer r.Release()
		f()
		return nil
	default:
		return fmt.Errorf("rate limit exceeded")
	}
}
