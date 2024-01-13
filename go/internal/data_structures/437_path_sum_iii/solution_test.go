package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

var tsc = []struct {
	root   []int
	target int
	ans    int
}{
	{
		[]int{10, 5, -3, 3, 2, tree_node.NULL, 11, 3, -2, tree_node.NULL, 1},
		8,
		3,
	},
	{
		[]int{5, 4, 8, 11, tree_node.NULL, 13, 4, 7, 2, tree_node.NULL, tree_node.NULL, 5, 1},
		22,
		3,
	},
}

func Test_pathSum_1(t *testing.T) {
	for _, tc := range tsc {
		root := tree_node.Ints2TreeNode(tc.root)
		result := Solution(root, tc.target)

		assert.Equal(t, tc.ans, result)
	}
}
