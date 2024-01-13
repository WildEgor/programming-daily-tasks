package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
	"testing"
)

func Test_maxDepth_1(t *testing.T) {

	root := tree_node.NewTreeNode(3)
	root.Insert(9)
	root.Insert(20)

	// Cause wtf why 9 left leaf fot 3 in binary tree?
	root.Left = tree_node.NewTreeNode(9)

	result := Solution(root)

	assert.Equal(t, 3, result)
}
