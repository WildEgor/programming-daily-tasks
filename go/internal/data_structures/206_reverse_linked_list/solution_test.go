package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_example_1(t *testing.T) {
	heads := []int{1, 2, 3, 4, 5}
	linked := NewListNode()

	for _, val := range heads {
		linked.Insert(val)
	}

	result := Solution(linked)

	assert.Equal(t, []int{5, 4, 3, 2, 1}, result.List())
}
