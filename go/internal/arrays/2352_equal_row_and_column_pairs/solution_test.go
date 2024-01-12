package equal_row_and_column_pairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_equalPairs_1(t *testing.T) {
	result := Solution([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}})
	assert.Equal(t, 1, result)
}

func Test_equalPairs_2(t *testing.T) {
	result := Solution([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})
	assert.Equal(t, 3, result)
}
