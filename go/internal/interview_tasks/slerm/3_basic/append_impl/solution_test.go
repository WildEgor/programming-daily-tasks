package append_impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_appendTo(t *testing.T) {
	result := appendTo([]int{1, 2, 3}, 4, 5)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}
