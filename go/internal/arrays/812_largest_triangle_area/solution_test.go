package largest_triangle_area

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	en   bool
	dots [][]int
	ans  float64
}{
	{
		false,
		[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
		2.0000,
	},
	{
		true,
		[][]int{{1, 0}, {0, 0}, {0, 1}},
		0.5000,
	},
}

func Test_removeDuplicates(t *testing.T) {
	for _, tc := range tsc {
		if tc.en {
			assert.Equal(t, tc.ans, Solution(tc.dots))
		}
	}
}
