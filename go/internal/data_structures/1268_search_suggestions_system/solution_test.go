package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	products   []string
	searchWord string
	result     [][]string
}{
	{
		products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
		searchWord: "mouse",
		result: [][]string{
			{"mobile", "moneypot", "monitor"},
		},
	},
	{
		products:   []string{"havana"},
		searchWord: "havana",
		result: [][]string{
			{"havana"},
		},
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tsc {
		result := Solution(tc.products, tc.searchWord)
		assert.Equal(t, tc.result, result)
	}
}
