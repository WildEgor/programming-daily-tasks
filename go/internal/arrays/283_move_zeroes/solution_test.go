package arrays_solutions_move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes_1(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	Solution(nums)

	assert.ElementsMatch(t, []int{1, 3, 12, 0, 0}, nums)
}

func Test_moveZeroes_2(t *testing.T) {
	nums := []int{0}
	Solution(nums)

	assert.ElementsMatch(t, []int{0}, nums)
}
