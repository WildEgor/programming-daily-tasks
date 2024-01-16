package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	a, b, c int
	exp     int
}{
	{
		2, 6, 5,
		3,
	},
	{
		4, 2, 7,
		1,
	},
	{
		1, 2, 3,
		0,
	},
}

func Test_minFlips_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.exp, Solution(tc.a, tc.b, tc.c))
	}
}
