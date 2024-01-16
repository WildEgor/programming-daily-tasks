package heap_impl

import (
	"cmp"
	"container/heap"
)

type Heap[T cmp.Ordered] struct {
	items []T
}

func NewHeap[T cmp.Ordered](items []T) Heap[T] {
	return Heap[T]{
		items: items,
	}
}

func (h *Heap[T]) GetMin() T {
	return heap.Pop(h)
}

func (h *Heap[T]) Len() int {
	return len(h.items)
}

func (h *Heap[T]) Less(i int, j int) bool {
	return h.items[i] < h.items[j]
}

func (h *Heap[T]) Push(v any) {
	h.items = append(h.items, v.(T))
}

func (h *Heap[T]) Swap(i int, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap[T]) Pop() any {
	lindx := len(h.items) - 1
	last := h.items[lindx]
	h.items = h.items[:lindx]
	return last
}

func (h *Heap[T]) Top() any {
	return h.items[0]
}

func (h *Heap[T]) Last() any {
	return h.items[len(h.items)-1]
}
