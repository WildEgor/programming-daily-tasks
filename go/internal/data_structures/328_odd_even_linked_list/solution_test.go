package odd_even_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"
	"testing"
)

func Test_oddEvenList_1(t *testing.T) {
	list := list_node.NewListNode(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	result := Solution(list)
	actual := result.List()
	assert.Equal(t, []int{1, 3, 5, 2, 4}, actual)
}
