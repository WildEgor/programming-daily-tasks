package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestSubarray_1(t *testing.T) {
	result := Solution([]int{1, 1, 0, 1})
	assert.Equal(t, 3, result)
}

func Test_longestSubarray_2(t *testing.T) {
	result := Solution([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})
	assert.Equal(t, 5, result)
}

func Test_longestSubarray_3(t *testing.T) {
	result := Solution([]int{1, 1, 1})
	assert.Equal(t, 2, result)
}

func Test_longestSubarray_4(t *testing.T) {
	result := Solution([]int{0, 0, 0})
	assert.Equal(t, 0, result)
}

func Test_longestSubarray_5(t *testing.T) {
	result := Solution([]int{1, 0, 0, 0, 0})
	assert.Equal(t, 1, result)
}
