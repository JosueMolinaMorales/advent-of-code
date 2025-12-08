package dsa

import "container/heap"

// A generic Heap type that works with any type T.
type Heap[T any] struct {
	data []T
	// The less function defines the ordering. For a min-heap, a < b is used.
	less func(a, b T) bool
}

// NewGenericMinHeap is a constructor function to create a new min-heap.
func NewMinHeap[T any](less func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{
		less: less,
	}
	// The container/heap functions expect a pointer to a struct that
	// satisfies the heap.Interface. The methods below implement this interface.
	heap.Init(h)
	return h
}

// Below are the five methods required by the sort.Interface and heap.Interface
func (h Heap[T]) Len() int           { return len(h.data) }
func (h Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

// Push and Pop use pointer receivers as they modify the slice length.
func (h *Heap[T]) Push(x any) {
	h.data = append(h.data, x.(T))
}

func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

// Helper methods for convenience
func (h *Heap[T]) PushItem(item T) {
	heap.Push(h, item)
}

func (h *Heap[T]) PopItem() T {
	return heap.Pop(h).(T)
}

func (h *Heap[T]) Peek() T {
	return h.data[0]
}
