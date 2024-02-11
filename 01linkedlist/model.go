package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(data int) {
	// create a new node, which will have data
	// and Next node address, initially there is no
	// Next node so Next will be nil
	newNode := &Node{Data: data, Next: nil}

	// initially if there is only one node then
	// Head should point to 1st node
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	// find lastNode, initially lastNode will be Head
	// and then check if Next is not nil in lastNode
	// then update the lastNode address
	lastNode := ll.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	// finally update the Node address part
	lastNode.Next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.Head

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}