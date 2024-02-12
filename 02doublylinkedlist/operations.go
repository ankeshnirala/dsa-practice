package doublylinkedlist

import "fmt"

func (dll *DLinkedList) Append(data int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}

	// if dll is empty set both head and tail as newNode
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	// Set the new node's prev pointer to the current tail
	newNode.Prev = dll.Tail
	// Set the current tail's next pointer to the new node
	dll.Tail.Next = newNode
	// Update the tail to be the new node
	dll.Tail = newNode
}

func (ll *DLinkedList) Displey() {
	current := ll.Head
	fmt.Println(&current)
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}
