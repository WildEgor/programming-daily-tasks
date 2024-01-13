package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"
	"testing"
)

func Test_example_1(t *testing.T) {
	heads := []int{2, 3, 4, 5}
	linked := list_node.NewListNode(1)

	for _, val := range heads {
		linked.Insert(val)
	}

	result := Solution(linked)

	assert.Equal(t, []int{5, 4, 3, 2, 1}, result.List())
}
