package template

import (
	"testing"
)

var tcs = []struct {
	cost   []int
	result int
}{
	{
		cost:   []int{10, 15, 20},
		result: 15,
	},
	{
		cost:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		result: 6,
	},
}

func Test_minCostClimbingStairs_1(t *testing.T) {
	for _, tc := range tcs {
		result := minCostClimbingStairs(tc.cost)
		if tc.result != result {
			t.Errorf("minCostClimbingStairs(%v) should have been %v, but was %v", tc.cost, tc.result, result)
		}
	}
}
