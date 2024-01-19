package remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	val  int
	k    int
}{
	{
		[]int{3, 2, 2, 3},
		3,
		2,
	},
}

func Test_removeElement_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.k, Solution(tc.nums, tc.val))
	}
}
