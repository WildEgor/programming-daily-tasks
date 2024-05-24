package buf_ch

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Pool(t *testing.T) {
	p := NewPool(func() any {
		return 5
	}, 3)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				p.Run(func(v any) {
					fmt.Println(v, i+v.(int))
				})
			}
		}()
	}
}
