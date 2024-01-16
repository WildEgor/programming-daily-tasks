package reorder_routes_to_make_all_paths_lead_to_the_city_zero

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	n      int
	conn   [][]int
	result int
}{
	{
		5,
		[][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
		2,
	},
}

func Test_minReorder_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.result, Solution(tc.n, tc.conn))
	}
}
