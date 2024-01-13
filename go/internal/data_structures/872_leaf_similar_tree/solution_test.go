package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

var tcs = []struct {
	root1 []int
	root2 []int
	ans   bool
}{
	{
		[]int{3, 5, 1, 6, 2, 9, 8, tree_node.NULL, tree_node.NULL, 7, 4},
		[]int{3, 5, 1, 6, 7, 4, 2, tree_node.NULL, tree_node.NULL, tree_node.NULL, tree_node.NULL, tree_node.NULL, tree_node.NULL, 9, 8},
		true,
	},
}

func Test_leafSimilar_1(t *testing.T) {

	root1 := tree_node.Ints2TreeNode(tcs[0].root1)
	root2 := tree_node.Ints2TreeNode(tcs[0].root2)

	result := Solution(root1, root2)

	assert.Equal(t, true, result)
}
