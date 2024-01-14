package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

// BFS
func maxDepth(root *tree_node.TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func maxDepthAlt(root *tree_node.TreeNode) int {
	queue := []*tree_node.TreeNode{root}

	var depth int

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return depth
}

var Solution = maxDepthAlt
