package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/list_node"

/**
In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.
For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
The twin sum is defined as the sum of a node and its twin.
Given the head of a linked list with even length, return the maximum twin sum of the linked list.
@list https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list
*/

func pairSum(head *list_node.ListNode) int {

	// stack := head.List()
	var stack []int

	curr := head
	for curr != nil {
		stack = append(stack, curr.Val)
		curr = curr.Next
	}

	result := 0
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	for i := 0; i < len(stack); i++ {
		sum := stack[i] + stack[len(stack)-i-1]
		result = max(result, sum)
	}

	return result
}

func pairSumAlt(head *list_node.ListNode) int {
	s, f := head, head
	ms := 0

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	f = s
	var prev *list_node.ListNode = nil
	for s != nil {
		next := s.Next
		s.Next = prev
		prev = s
		s = next
	}

	for head != nil && head != f {
		sum := head.Val + prev.Val
		if sum > ms {
			ms = sum
		}
		head = head.Next
		prev = prev.Next
	}

	return ms
}

var Solution = pairSumAlt
