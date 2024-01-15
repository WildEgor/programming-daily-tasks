package domino_and_tromino_tiling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	n      int
	result int
}{
	{
		3,
		5,
	},
	{
		1,
		1,
	},
}

func Test_numTilings_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.result, Solution(tc.n))
	}
}
