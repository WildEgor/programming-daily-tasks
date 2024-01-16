package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	text1  string
	text2  string
	result int
}{
	{
		"abcde",
		"ace",
		3,
	},
	{
		"abc",
		"abc",
		3,
	},
	{
		"abc",
		"def",
		0,
	},
}

func Test_longestCommonSubsequence_1(t *testing.T) {
	for _, ts := range tsc {
		assert.Equal(t, ts.result, Solution(ts.text1, ts.text2))
	}
}
