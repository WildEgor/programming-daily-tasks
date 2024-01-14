package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
You are given the root of a binary search tree (BST) and an integer val.
Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.
@link https://leetcode.com/problems/search-in-a-binary-search-tree/
*/

func searchBST(root *tree_node.TreeNode, val int) *tree_node.TreeNode {
	if root == nil {
		return nil
	}

	// Why it effectively works? If we swap the order of the conditions, it will work slower...
	if val == root.Val {
		return root
	}

	if val > root.Val {
		return searchBST(root.Right, val) // move to right subtree
	}

	return searchBST(root.Left, val) // else move to left subtree
}

var Solution = searchBST
