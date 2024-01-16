package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	word1  string
	word2  string
	result int
}{
	{
		"horse",
		"ros",
		3,
	},
	{
		"intention",
		"execution",
		5,
	},
}

func Test_minDistance_1(t *testing.T) {
	for _, tc := range tsc {
		assert.Equal(t, tc.result, Solution(tc.word1, tc.word2))
	}
}
