package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

var tcs = []struct {
	root []int
	val  int
	ans  []int
}{
	{
		[]int{4, 2, 7, 1, 3},
		2,
		[]int{2, 1, 3},
	},
	{
		[]int{4, 2, 7, 1, 3},
		5,
		[]int{},
	},
}

func Test_searchBST_1(t *testing.T) {
	for _, tc := range tcs {
		result := Solution(tree_node.Ints2TreeNode(tc.root), tc.val)
		assert.Equal(t, tc.ans, result.ToList())
	}
}
