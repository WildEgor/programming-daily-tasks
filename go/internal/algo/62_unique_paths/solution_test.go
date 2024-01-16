package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	m      int
	n      int
	result int
}{
	{
		3,
		7,
		28,
	},
	{
		3,
		2,
		3,
	},
}

func Test_unique_paths_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.result, Solution(tc.m, tc.n))
	}
}
