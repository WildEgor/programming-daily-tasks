package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

var tsc = []struct {
	root []int
	p    int
	q    int
	ans  int
}{
	{
		[]int{3, 5, 1, 6, 2, 0, 8, tree_node.NULL, tree_node.NULL, 7, 4},
		5,
		1,
		3,
	},
	{
		[]int{3, 5, 1, 6, 2, 0, 8, tree_node.NULL, tree_node.NULL, 7, 4},
		5,
		4,
		5,
	},
	{
		[]int{1, 2},
		1,
		2,
		1,
	},
}

func Test_lowestCommonAncestor_1(t *testing.T) {

	for _, tc := range tsc {
		root := tree_node.Ints2TreeNode(tc.root)
		q := tree_node.NewTreeNode(tc.q)
		p := tree_node.NewTreeNode(tc.p)
		result := Solution(root, p, q)
		assert.Equal(t, tc.ans, result.Val)
	}
}
