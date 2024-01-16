package find_peak_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	input  []int
	output int
}{
	{[]int{1, 2, 3, 1}, 2},
	{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
}

func Test_findPeakElement_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.output, Solution(tc.input))
	}
}
