package model

type MyQueue[T any] struct {
	data []T
}

func (q *MyQueue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}

	return q.data[0], true
}

func (q *MyQueue[T]) Size() int {
	return len(q.data)
}

// IsEmpty checks if the queue is empty
func (q *MyQueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q MyQueue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.data[0], true
}

func (q *MyQueue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

func NewMyQueue[T any]() *MyQueue[T] {
	return &MyQueue[T]{
		data: make([]T, 0),
	}
}
