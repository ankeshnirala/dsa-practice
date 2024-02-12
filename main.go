package main

import (
	linkedlist "dsa-practice/01linkedlist"
	doublylinkedlist "dsa-practice/02doublylinkedlist"
	stack "dsa-practice/03stack"
	"fmt"
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

	// dll.Displey()

	stack := &stack.Stack{}
	// stack.Push("10")
	// stack.Push("20")
	// stack.Push("30")

	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Peek())

	// fmt.Println(stack.BalancedParenthesis("({)}"))
	fmt.Println(stack.ReverseString("ankesh"))
	// stack.Display()
}
