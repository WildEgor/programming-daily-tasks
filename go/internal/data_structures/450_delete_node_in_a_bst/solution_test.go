package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

var tcs = []struct {
	root []int
	key  int
	ans  []int
}{
	{
		[]int{5, 3, 6, 2, 4, tree_node.NULL, 7},
		3,
		[]int{5, 4, 6, 2, tree_node.NULL, 7},
	},
	{
		[]int{5, 3, 6, 2, 4, tree_node.NULL, 1},
		0,
		[]int{5, 3, 6, 2, 4, tree_node.NULL, 1},
	},
}

func Test_longestZigZag_1(t *testing.T) {
	for _, tc := range tcs {
		result := Solution(tree_node.Ints2TreeNode(tc.root), tc.key)
		assert.Equal(t, tc.ans, result.ToList())
	}
}
