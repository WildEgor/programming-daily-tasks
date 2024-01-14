package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

var tcs = []struct {
	root []int
	ans  []int
}{
	{
		[]int{1, 2, 3, tree_node.NULL, 5, tree_node.NULL, 4},
		[]int{1, 3, 4},
	},
	{
		[]int{1, tree_node.NULL, 3},
		[]int{1, 3},
	},
	{
		[]int{},
		[]int{},
	},
}

func Test_rightSideView_1(t *testing.T) {
	for _, tc := range tcs {
		result := Solution(tree_node.Ints2TreeNode(tc.root))
		assert.Equal(t, tc.ans, result)
	}
}
