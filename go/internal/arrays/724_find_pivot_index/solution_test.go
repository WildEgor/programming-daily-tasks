package find_pivot_index

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotIndex_1(t *testing.T) {
	result := Solution([]int{1, 7, 3, 6, 5, 6})
	assert.Equal(t, 3, result)
}

func Test_pivotIndex_2(t *testing.T) {
	result := Solution([]int{1, 2, 3})
	assert.Equal(t, -1, result)
}

func Test_pivotIndex_3(t *testing.T) {
	result := Solution([]int{2, 1, -1})
	assert.Equal(t, 0, result)
}
