package letter_combinations_of_a_phone_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	digits string
	result []string
}{
	{
		"23",
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
	{
		"",
		[]string{},
	},
	{
		"2",
		[]string{"a", "b", "c"},
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tcs {
		result := Solution(tc.digits)
		assert.Equal(t, tc.result, result)
	}
}
