package linkedlistrevision01

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Insert(data string) {

	// create a new node
	newNode := &Node{Data: data, Next: nil}

	// if head == nil, then head = newNode
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	// initially lastNode = head
	lastNode := ll.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}
