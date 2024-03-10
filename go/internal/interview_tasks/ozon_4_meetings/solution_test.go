package ozon_4_meetings

import (
	"reflect"
	"testing"
)

var tsc = []struct {
	input  [][]int
	output [][]int
}{
	{
		input:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		output: [][]int{{1, 6}, {8, 10}, {15, 18}},
	},
	{
		input:  [][]int{{1, 4}, {4, 5}},
		output: [][]int{{1, 5}},
	},
}

func Test_merge(t *testing.T) {
	for _, test := range tsc {
		result := merge(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("expected: %v, got: %v", test.output, result)
		}
	}
}
