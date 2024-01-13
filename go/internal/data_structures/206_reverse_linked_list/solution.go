package reverse_linked_list

import "github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"

/**
Given the head of a singly linked list, reverse the list, and return the reversed list.
@link https://leetcode.com/problems/reverse-linked-list
*/

func reverseList(head *list_node.ListNode) (prev *list_node.ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}

var Solution = reverseList
