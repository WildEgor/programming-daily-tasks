package odd_even_linked_list

import "github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"

/**
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
The first node is considered odd, and the second node is even, and so on.
Note that the relative order inside both the even and odd groups should remain as it was in the input.
You must solve the problem in O(1) extra space complexity and O(n) time complexity.
@link https://leetcode.com/problems/odd-even-linked-list
*/

func oddEvenList(head *list_node.ListNode) *list_node.ListNode {
	if head == nil {
		return nil
	}

	// Save first odd position and next as even
	oddNode := head
	evenNode := head.Next
	evenHead := evenNode

	// Check if head have next (even) and another odd nodes
	for evenNode != nil && evenNode.Next != nil {
		oddNode.Next = oddNode.Next.Next
		oddNode = oddNode.Next
		evenNode.Next = evenNode.Next.Next
		evenNode = evenNode.Next
	}

	oddNode.Next = evenHead

	return head
}

var Solution = oddEvenList
