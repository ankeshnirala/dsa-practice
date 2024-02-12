package doublylinkedlist

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DLinkedList struct {
	Head *Node
	Tail *Node
}
