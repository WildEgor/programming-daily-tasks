package nearest_exit_from_entrance_in_maze

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	maze [][]byte
	ent  []int
	r    int
}{
	{
		[][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}},
		[]int{1, 2},
		1,
	},
}

func Test_nearestExit_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.r, Solution(tc.maze, tc.ent))
	}
}
