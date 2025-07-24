package arraysandslices

type Int interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueue[T Int] struct {
	values []T
	filled int
	head   int
	tail   int
}

func NewCircularQueue[T Int](size int) CircularQueue[T] {
	return CircularQueue[T]{
		values: make([]T, size),
		head:   0,
		filled: 0,
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	q.values[q.tail] = value

	q.tail++
	if q.tail >= len(q.values) {
		q.tail = 0
	}

	q.filled++

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	q.values[q.head] = 0
	q.head++
	q.filled--

	if q.head >= len(q.values) {
		q.head = 0
	}

	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.Empty() {
		return -1
	}

	return q.values[q.head]
}

func (q *CircularQueue[T]) Back() T {
	if q.Empty() {
		return -1
	}

	idx := q.tail - 1

	if idx < 0 {
		idx = len(q.values) - 1
	}

	return q.values[idx]
}

func (q *CircularQueue[T]) Empty() bool {
	return q.filled == 0
}

func (q *CircularQueue[T]) Full() bool {
	return q.filled == len(q.values)
}
