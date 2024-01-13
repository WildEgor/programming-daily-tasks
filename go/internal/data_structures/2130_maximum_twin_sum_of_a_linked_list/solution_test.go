package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"
	"testing"
)

func Test_pairSum_1(t *testing.T) {
	list := list_node.NewListNode(5)
	list.Insert(4)
	list.Insert(2)
	list.Insert(1)

	result := Solution(list)

	assert.Equal(t, 6, result)
}

func Test_pairSum_2(t *testing.T) {
	list := list_node.NewListNode(4)
	list.Insert(2)
	list.Insert(2)
	list.Insert(3)

	result := Solution(list)

	assert.Equal(t, 7, result)
}
