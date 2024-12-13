package linkedlist

import "testing"

func TestReverse(t *testing.T) {
	list := &LinkedList{}
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	list.Reverse()

	if list.Head.Value != 3 {
		t.Error("Expected 3, got", list.Head.Value)
	}
	if list.Head.Next.Value != 2 {
		t.Error("Expected 2, got", list.Head.Next.Value)
	}
	if list.Head.Next.Next.Value != 1 {
		t.Error("Expected 1, got", list.Head.Next.Next.Value)
	}
	if list.Head.Next.Next.Next != nil {
		t.Error("Expected nil, got", list.Head.Next.Next.Next)
	}
}
