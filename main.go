package main

import (
	linkedlist "dsa-practice/01linkedlist"
	doublylinkedlist "dsa-practice/02doublylinkedlist"
)

func main() {
	ll := &linkedlist.LinkedList{}

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	// ll.Display()

	dll := &doublylinkedlist.DLinkedList{}
	dll.Append(10)
	dll.Append(20)
	dll.Append(30)
	// dll.Append(40)

	dll.Displey()
}
