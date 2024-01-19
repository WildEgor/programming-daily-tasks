package remove_duplicates_from_sorted_array_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	nums []int
	r    int
}{
	{
		[]int{1, 1, 1, 2, 2, 3},
		5,
	},
	{
		[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		7,
	},
}

func Test_removeDuplicates_1(t *testing.T) {
	for _, tc := range tsc {
		assert.Equal(t, tc.r, Solution(tc.nums))
	}
}
