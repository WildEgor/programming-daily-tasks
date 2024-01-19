package rotate_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	k    int
	r    []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 6, 7, 1, 2, 3, 4},
	},
}

func Test_rotate_1(t *testing.T) {
	for _, tc := range tcs {
		Solution(tc.nums, tc.k)

		assert.Equal(t, tc.r, tc.nums)
	}
}
