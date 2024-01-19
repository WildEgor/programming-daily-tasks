package remove_duplicates_from_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	r    int
}{
	{
		[]int{1, 1, 2},
		2,
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.r, Solution(tc.nums))
	}
}
