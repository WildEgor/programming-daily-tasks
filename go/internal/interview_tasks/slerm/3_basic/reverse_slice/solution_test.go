package reverse_slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverse(tt.arr); !assert.Equal(t, got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
