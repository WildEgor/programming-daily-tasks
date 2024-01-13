package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
Return the number of good nodes in the binary tree.
@link https://leetcode.com/problems/count-good-nodes-in-binary-tree/
*/

func goodNodes(root *tree_node.TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *tree_node.TreeNode, max int) int {
	// Base case
	if root == nil {
		return 0
	}

	// Try to find max
	if root.Val >= max {
		return 1 + dfs(root.Left, root.Val) + dfs(root.Right, root.Val)
	}

	// If not found, continue search
	return dfs(root.Left, max) + dfs(root.Right, max)
}

var Solution = goodNodes
