package stack

type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{elements: make([]interface{}, 0)}
}

func (s *Stack) Push(e interface{}) {
	s.elements = append(s.elements, e)
}

// Pop removes the last element in the stack and returns it
func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e
}

// Peek returns the last element in the stack without removing it
func (s *Stack) Peek() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Empty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Clear() {
	s.elements = make([]interface{}, 0)
}

func (s *Stack) Values() []interface{} {
	return s.elements
}
