package template

import (
	"github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"
)

/**
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.
The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).
@link https://leetcode.com/problems/path-sum-iii/
*/

func pathSum(root *tree_node.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	m := make(map[int]int)
	m[0] = 1

	var result int
	sumPaths(&result, root, 0, targetSum, m)

	return result
}

func sumPaths(res *int, node *tree_node.TreeNode, s int, targetSum int, cache map[int]int) {
	if node == nil {
		return
	}

	s += node.Val
	if counter, ok := cache[s-targetSum]; ok {
		*res += counter
	}

	if counter, ok := cache[s]; ok {
		cache[s] = counter + 1
	} else {
		cache[s] = 1
	}

	sumPaths(res, node.Left, s, targetSum, cache)
	sumPaths(res, node.Right, s, targetSum, cache)

	cache[s]--
}

var Solution = pathSum
