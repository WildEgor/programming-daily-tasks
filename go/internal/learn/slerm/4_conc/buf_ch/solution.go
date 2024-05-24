package buf_ch

import (
	"errors"
	"time"
)

type Factory func() any

type Processor func(any)

type IPool interface {
	Run(Processor)
	RunWithTimeout(Processor, time.Duration) error
}

var _ IPool = (*poolInner)(nil)

type poolInner struct {
	items chan any
}

func NewPool(f Factory, count int) IPool {
	pl := &poolInner{
		items: make(chan any, count),
	}
	for i := 0; i < count; i++ {
		pl.items <- f()
	}
	return pl
}

// Run receive items, call Processor and return item back
func (pi *poolInner) Run(p Processor) {
	item := <-pi.items
	defer func() {
		pi.items <- item
	}()
	p(item)
}

// RunWithTimeout wait while Processor work or return error
func (pi *poolInner) RunWithTimeout(p Processor, dur time.Duration) error {
	select {
	case item := <-pi.items:
		defer func() {
			pi.items <- item
		}()

		p(item)
	case <-time.After(dur):
		return errors.New("error: exceed timeout")
	}

	return nil
}
