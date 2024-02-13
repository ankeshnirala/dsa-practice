package main

import (
	linkedlist "dsa-practice/01linkedlist"
	doublylinkedlist "dsa-practice/02doublylinkedlist"
	stack "dsa-practice/03stack"
	searching "dsa-practice/05searching"
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
	stack.Push("10")
	stack.Push("20")
	stack.Push("30")

	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Peek())

	// fmt.Println(stack.BalancedParenthesis("({)}"))
	// fmt.Println(stack.ReverseString("ankesh"))
	stack.Display()

	// searching
	ls := searching.Input{2, 3, 3, 3, 5, 9, 4, 6, 1}
	fmt.Println(ls.LinearSearch(5))

	bs := searching.Input{2, 3, 5, 7, 9, 11}
	fmt.Println(bs.BinarySearch(5))

	fmt.Println(ls.CountTargetOccurence(3))
	fmt.Println(ls.FindMaxMin())
	fmt.Println(ls.TwoSum(14))

}
