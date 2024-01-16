package successful_pairs_of_spells_and_positions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	spells  []int
	potions []int
	success int64
	ans     []int
}{
	{
		[]int{5, 1, 3},
		[]int{1, 2, 3, 4, 5},
		7,
		[]int{4, 0, 3},
	},
	{
		[]int{3, 1, 2},
		[]int{8, 5, 8},
		16,
		[]int{2, 0, 2},
	},
}

func Test_successfulPairs_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.ans, Solution(tc.spells, tc.potions, tc.success))
	}
}
