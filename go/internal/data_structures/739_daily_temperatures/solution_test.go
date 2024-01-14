package template

import (
	"testing"
)

var tsc = []struct {
	temperatures []int
	answer       []int
}{
	{
		[]int{73, 74, 75, 71, 69, 72, 76, 73},
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tsc {
		got := dailyTemperatures(tc.temperatures)
		for i, v := range got {
			if v != tc.answer[i] {
				t.Errorf("expected %v, got %v", tc.answer[i], v)
			}
		}
	}
}
