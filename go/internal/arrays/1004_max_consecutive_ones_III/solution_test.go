package max_consecutive_ones_III

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestOnes_1(t *testing.T) {
	result := Solution([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 2)
	assert.Equal(t, 6, result)
}

func Test_longestOnes_2(t *testing.T) {
	result := Solution([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	assert.Equal(t, 10, result)
}
