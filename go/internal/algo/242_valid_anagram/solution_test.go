package valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	s string
	t string
	r bool
}{
	{
		"anagram",
		"nagaram",
		true,
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.r, Solution(tc.s, tc.t))
	}
}
