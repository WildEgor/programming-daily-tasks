package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	input  int
	output int
}{
	{
		input:  4,
		output: 4,
	},
	{
		input:  25,
		output: 1389537,
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tsc {
		assert.Equal(t, tc.output, Solution(tc.input))
	}
}
