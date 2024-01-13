package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
You are given the root of a binary tree.
A ZigZag path for a binary tree is defined as follow:
Choose any node in the binary tree and a direction (right or left).
If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
Change the direction from right to left or from left to right.
Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
Return the longest ZigZag path contained in that tree.
@link https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
*/

func longestZigZag(root *tree_node.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func dfs(root *tree_node.TreeNode, left bool, length int) int {
	if root == nil {
		return length
	}

	if left {
		return max(dfs(root.Left, true, 0), dfs(root.Right, false, length+1))
	}

	return max(dfs(root.Left, true, length+1), dfs(root.Right, false, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Solution = longestZigZag
