package most_common_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	paragraph string
	banned    []string
	expected  string
}{
	{
		"Bob hit a ball, the hit BALL flew far after it was hit.",
		[]string{"hit"},
		"ball",
	},
}

func Test_mostCommonWord(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.expected, mostCommonWord(tc.paragraph, tc.banned))
	}
}
