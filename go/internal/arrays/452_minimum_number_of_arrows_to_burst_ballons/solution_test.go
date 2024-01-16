package minimum_number_of_arrows_to_burst_ballons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	input  [][]int
	output int
}{
	{
		input:  [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
		output: 2,
	},
	{
		input:  [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		output: 4,
	},
	{
		input:  [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		output: 2,
	},
}

func Test_findMinArrowShots_1(t *testing.T) {
	for _, tc := range tsc {
		assert.Equal(t, tc.output, Solution(tc.input))
	}
}
