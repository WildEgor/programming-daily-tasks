package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q
as descendants (where we allow a node to be a descendant of itself).”
@link https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
*/

func lowestCommonAncestor(root, p, q *tree_node.TreeNode) *tree_node.TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == p && root.Right == q {
		return root
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

var Solution = lowestCommonAncestor
