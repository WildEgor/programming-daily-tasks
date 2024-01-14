package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	n        int
	expected []int
}{
	{
		n:        2,
		expected: []int{0, 1, 1},
	},
	{
		n:        5,
		expected: []int{0, 1, 1, 2, 1, 2},
	},
}

func Test_countBits_1(t *testing.T) {
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, Solution(tc.n))
	}
}
