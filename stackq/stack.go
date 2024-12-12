package stackq

type Stack[T Number] struct {
	elements []T
}

// Push добавляет элемент в стек
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop удаляет верхний элемент из стека и возвращает его
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

// Peek возвращает верхний элемент из стека без его удаления
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}
