package main

import linkedlist "dsa-practice/01linkedlist"

func main() {
	ll := &linkedlist.LinkedList{}

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	ll.Display()
}