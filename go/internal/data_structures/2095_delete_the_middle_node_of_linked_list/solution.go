package template

/**
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.
For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
@link https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func (s *ListNode) Insert(val int) {
	st := &ListNode{
		Val:  val,
		Next: s.Next,
	}

	s.Next = st
}

func (s *ListNode) List() []int {
	var list []int
	var revesed []int

	list = append(list, s.Val)

	for st := s.Next; st != nil; st = st.Next {
		list = append(list, st.Val)
	}

	for i := len(list) - 1; i >= 0; i-- {
		revesed = append(revesed, list[i])
	}

	return revesed
}

func (s *ListNode) Size() int {
	size := 0
	for st := s.Next; st != nil; st = st.Next {
		size++
	}

	return size
}

func (s *ListNode) DeleteMiddle() *ListNode {
	var prev *ListNode
	i := 0

	var head *ListNode

	index := int(s.Size() / 2)

	for st := s.Next; st != nil; st = st.Next {
		if head == nil {
			head = st
		}

		if index == i && prev != nil {
			prev.Next = st.Next
			break
		}

		prev = st
		i++
	}

	return head
}

func deleteMiddle(head *ListNode) *ListNode {
	return head.DeleteMiddle()
}

func deleteMiddleAlt(head *ListNode) *ListNode {
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
