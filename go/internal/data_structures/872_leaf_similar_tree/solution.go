package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
Two binary trees are considered leaf-similar if their leaf value sequence is the same.
Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.
@link https://leetcode.com/problems/leaf-similar-trees
*/

func leafSimilar(root1 *tree_node.TreeNode, root2 *tree_node.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	var leaves1, leaves2 []int

	dfs(root1, &leaves1)
	dfs(root2, &leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}

func dfs(root *tree_node.TreeNode, leaves *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	}

	dfs(root.Left, leaves)
	dfs(root.Right, leaves)
}

var Solution = leafSimilar
