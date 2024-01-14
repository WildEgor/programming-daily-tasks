package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	intervals [][]int
	output    int
}{
	{
		intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		output:    1,
	},
	{
		intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
		output:    2,
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tsc {
		assert.Equal(t, tc.output, eraseOverlapIntervals(tc.intervals))
	}
}
