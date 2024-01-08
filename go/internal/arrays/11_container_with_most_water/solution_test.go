package arrays_11_container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes_1(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := Solution(nums)

	assert.Equal(t, 49, result)
}

func Test_moveZeroes_2(t *testing.T) {
	nums := []int{1, 1}
	result := Solution(nums)

	assert.Equal(t, 1, result)
}
