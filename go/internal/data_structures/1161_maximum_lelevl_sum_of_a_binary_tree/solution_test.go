package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

var tcs = []struct {
	root []int
	ans  int
}{
	{
		[]int{1, 7, 0, 7, -8, tree_node.NULL, tree_node.NULL},
		2,
	},
	{
		[]int{989, tree_node.NULL, 10250, 98693, -89388, tree_node.NULL, tree_node.NULL, tree_node.NULL, -32127},
		2,
	},
}

func Test_maxLevelSum_1(t *testing.T) {
	for _, tc := range tcs {
		result := Solution(tree_node.Ints2TreeNode(tc.root))
		assert.Equal(t, tc.ans, result)
	}
}
