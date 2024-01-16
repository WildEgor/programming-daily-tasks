package nearest_exit_from_entrance_in_maze

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	grid [][]int
	r    int
}{
	{
		[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
		4,
	},
}

func Test_nearestExit_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.r, Solution(tc.grid))
	}
}
