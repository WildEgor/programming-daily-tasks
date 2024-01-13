package list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(head int) *ListNode {
	return &ListNode{
		Val: head,
	}
}

func (s *ListNode) Size() int {
	size := 0
	for st := s.Next; st != nil; st = st.Next {
		size++
	}

	return size
}

func (s *ListNode) Insert(val int) {
	current := s
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (s *ListNode) List() []int {
	var list []int

	curr := s
	for curr != nil {
		list = append(list, curr.Val)
		curr = curr.Next
	}

	return list
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
