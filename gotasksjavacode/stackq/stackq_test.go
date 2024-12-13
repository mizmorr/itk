package stackq

import "testing"

func TestStack(t *testing.T) {
	s := &Stack[float64]{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if val, _ := s.Peek(); val != 3 {
		t.Error("Expected 3")
	}

	popped, _ := s.Pop()
	if popped != 3 {
		t.Error("Expected 3, got", popped)
	}
}

func TestQueue(t *testing.T) {
	q := &Queue[uint64]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if val, _ := q.Peek(); val != 1 {
		t.Error("Expected 1")
	}

	popped, _ := q.Dequeue()
	if popped != 1 {
		t.Error("Expected 1, got", popped)
	}
}
