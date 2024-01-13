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
		[]int{3, 1, 4, 3, tree_node.NULL, 1, 5},
		4,
	},
	{
		[]int{3, 3, tree_node.NULL, 4, 2},
		3,
	},
	{
		[]int{1},
		1,
	},
}

func Test_goodNodes_1(t *testing.T) {

	for _, tc := range tcs {
		root := tree_node.Ints2TreeNode(tc.root)
		result := Solution(root)

		assert.Equal(t, tc.ans, result)
	}
}
