package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	isConnected [][]int
	ans         int
}{
	{
		[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
		2,
	},
	{
		[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		3,
	},
}

func Test_findCircleNum_1(t *testing.T) {
	for _, tc := range tcs {
		result := Solution(tc.isConnected)
		assert.Equal(t, tc.ans, result)
	}
}
