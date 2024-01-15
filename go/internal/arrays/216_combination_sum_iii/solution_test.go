package combination_sum_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	k      int
	n      int
	result [][]int
}{
	{
		3,
		7,
		[][]int{
			{1, 2, 4},
		},
	},
	{
		3,
		9,
		[][]int{
			{1, 2, 6},
			{1, 3, 5},
			{2, 3, 4},
		},
	},
}

func Test_combinationSum3_1(t *testing.T) {
	for _, tc := range tcs {
		result := combinationSum3(tc.k, tc.n)
		assert.Equal(t, tc.result, result)
	}
}
