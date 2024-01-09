package find_the_highest_altitude

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestAltitude_1(t *testing.T) {
	result := Solution([]int{-5, 1, 5, 0, -7})
	assert.Equal(t, 1, result)
}

func Test_largestAltitude_2(t *testing.T) {
	result := Solution([]int{-4, -3, -2, -1, 4, 3, 2})
	assert.Equal(t, 0, result)
}
