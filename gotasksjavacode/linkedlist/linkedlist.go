package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// Добавление элемента в начало списка
func (list *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: list.Head}
	list.Head = newNode
}

// Печать всех элементов списка
func (list *LinkedList) Print() {
	for node := list.Head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println("nil")
}

func (list *LinkedList) Reverse() {
	var prev, current *Node = nil, list.Head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	list.Head = prev
}
