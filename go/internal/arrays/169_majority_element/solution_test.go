package majority_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	r    int
}{
	//{
	//	[]int{3, 2, 3},
	//	3,
	//},
	//{
	//	[]int{2, 2, 1, 1, 1, 2, 2},
	//	2,
	//},
	{
		[]int{4, 5, 4},
		4,
	},
}

func Test_majorityElement_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.r, Solution(tc.nums))
	}
}
