package kth_largest_element_in_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	input  []int
	k      int
	output int
}{
	{
		input:  []int{3, 2, 1, 5, 6, 4},
		k:      2,
		output: 5,
	},
}

func Test_rob_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.output, Solution(tc.input, tc.k))
	}
}
