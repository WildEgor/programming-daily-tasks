package reverse_linked_list

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

	list = append(list, s.Val)

	for st := s.Next; st != nil; st = st.Next {
		list = append(list, st.Val)
	}

	return list
}

// TODO: !!!!!!
func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}

var Solution = reverseList
