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
		input:  1,
		output: 1,
	},
}

func Test_solution(t *testing.T) {
	for _, test := range tsc {
		assert.Equal(t, test.output, solution(test.input))
	}
}
