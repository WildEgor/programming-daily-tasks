package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	input  []int
	output int
}{
	{
		input:  []int{1, 2, 3, 1},
		output: 4,
	},
	{
		input:  []int{1, 2},
		output: 2,
	},
}

func Test_rob_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.output, rob(tc.input))
	}
}
