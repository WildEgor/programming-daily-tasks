package template

import (
	"testing"
)

var testCases = []struct {
	name     string
	input    []int
	expected int
}{
	{
		name:     "example 1",
		input:    []int{2, 2, 1},
		expected: 1,
	},
	{
		name:     "example 2",
		input:    []int{4, 1, 2, 1, 2},
		expected: 4,
	},
	{
		name:     "example 3",
		input:    []int{1},
		expected: 1,
	},
}

func Test_singleNumber_1(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Solution(tc.input)
			if actual != tc.expected {
				t.Errorf("expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}
