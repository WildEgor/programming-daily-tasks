package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"
	"testing"
)

func Test_deleteMiddle_1(t *testing.T) {
	heads := []int{3, 4, 7, 1, 2, 6}
	linked := list_node.NewListNode(1)

	for _, val := range heads {
		linked.Insert(val)
	}

	head := Solution(linked)
	result := head.List()

	assert.ElementsMatch(t, []int{1, 3, 4, 1, 2, 6}, result)
}
