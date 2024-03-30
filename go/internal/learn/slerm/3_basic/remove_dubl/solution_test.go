package remove_dubl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeDuplicates(tt.arr); !assert.Equal(t, got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
