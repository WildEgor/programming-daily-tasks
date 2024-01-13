package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"

/**
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.
For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
@link https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list
*/

func deleteMiddleAlt(head *list_node.ListNode) *list_node.ListNode {
	current := head
	count := 0

	for current != nil {
		current = current.Next
		count++
	}

	if count == 1 {
		return nil
	}
	middle := count / 2

	preNode := head
	for i := 0; i < middle-1; i++ {
		preNode = preNode.Next
	}
	preNode.Next = preNode.Next.Next
	return head
}

var Solution = deleteMiddleAlt
