package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/tree_node"

/**
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
Basically, the deletion can be divided into two stages:
Search for a node to remove.
If the node is found, delete the node.
@link https://leetcode.com/problems/delete-node-in-a-bst/
*/

func deleteNode(root *tree_node.TreeNode, key int) *tree_node.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}

		if root.Left != nil && root.Right != nil {
			minSubTreeNode := findMin(root.Right)
			leftSubTree := root.Left
			rightSubTree := deleteNode(root.Right, minSubTreeNode.Val)
			minSubTreeNode.Left = leftSubTree
			minSubTreeNode.Right = rightSubTree
			return minSubTreeNode
		}
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func findMin(root *tree_node.TreeNode) *tree_node.TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return findMin(root.Left)
}

var Solution = deleteNode
