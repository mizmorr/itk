package stackq

type Queue[T Number] struct {
	elements []T
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue удаляет элемент из начала очереди и возвращает его
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	return q.elements[0], true
}
