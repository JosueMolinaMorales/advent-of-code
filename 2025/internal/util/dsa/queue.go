package dsa

// Queue is a generic FIFO data structure
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an item to the back of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue
// Returns the zero value of T if the queue is empty
func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		var zero T
		return zero
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Peek returns the front item without removing it
// Returns the zero value of T if the queue is empty
func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		var zero T
		return zero
	}
	return q.items[0]
}

// IsEmpty returns true if the queue has no items
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Clear removes all items from the queue
func (q *Queue[T]) Clear() {
	q.items = make([]T, 0)
}
