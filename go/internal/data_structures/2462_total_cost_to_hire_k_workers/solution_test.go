package total_cost_to_hire_k_workers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	costs      []int
	k          int
	candidates int
	result     int64
}{
	{
		[]int{17, 12, 10, 2, 7, 2, 11, 20, 8},
		3,
		4,
		11,
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.result, Solution(tc.costs, tc.k, tc.candidates))
	}
}
