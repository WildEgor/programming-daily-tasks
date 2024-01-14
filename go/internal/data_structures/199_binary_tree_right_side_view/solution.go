package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
@link https://leetcode.com/problems/binary-tree-right-side-view/
*/

func rightSideView(root *tree_node.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	dfs(root, 1, &result)
	return result
}

func dfs(root *tree_node.TreeNode, level int, result *[]int) {
	if root == nil {
		return
	}

	if len(*result) < level {
		*result = append(*result, root.Val)
	}

	dfs(root.Right, level+1, result)
	dfs(root.Left, level+1, result) // need for case when left subtree is deeper than right, for example [1,2]
}

var Solution = rightSideView
