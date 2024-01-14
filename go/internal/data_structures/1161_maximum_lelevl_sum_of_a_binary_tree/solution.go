package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
Return the smallest level x such that the sum of all the values of nodes at level x is maximal.
@link https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
*/

func maxLevelSum(root *tree_node.TreeNode) int {
	var levelSums []int
	dfs(root, 0, &levelSums)

	maxSum := levelSums[0]
	idx := 1
	for i := 1; i < len(levelSums); i++ {
		if levelSums[i] > maxSum {
			maxSum = levelSums[i]
			idx = i + 1
		}
	}

	return idx
}

func dfs(node *tree_node.TreeNode, level int, levelSums *[]int) {
	if node == nil {
		return
	}

	if len(*levelSums) == level {
		*levelSums = append(*levelSums, node.Val) // for root node
	} else {
		(*levelSums)[level] += node.Val
	}

	dfs(node.Left, level+1, levelSums)
	dfs(node.Right, level+1, levelSums)
}

var Solution = maxLevelSum
