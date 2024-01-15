package smallest_number_in_infinite_set

import (
	"container/heap"
	heap_impl "github.com/wildegor/daily_tasks/go/internal/data_structures/heap"
)

type SmallestInfiniteSetAlt struct {
	cursor int
	single map[int]bool
	heap   *heap_impl.Heap[int]
}

func NewSmallestInfiniteSetAlt() SmallestInfiniteSetAlt {
	return SmallestInfiniteSetAlt{
		cursor: 1,
		single: make(map[int]bool),
		heap:   &heap_impl.Heap[int]{},
	}
}

func (s *SmallestInfiniteSetAlt) PopSmallest() int {
	if s.heap.Len() > 0 {
		top := heap.Pop(s.heap).(int)
		s.single[top] = false
		return top
	}

	ret := s.cursor
	s.cursor++
	return ret
}

func (s *SmallestInfiniteSetAlt) AddBack(num int) {
	if num >= s.cursor {
		return
	}
	if s.single[num] {
		return
	}
	s.single[num] = true
	heap.Push(s.heap, num)
}
