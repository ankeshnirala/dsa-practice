package stack

import (
	"fmt"
)

func (s *Stack) Push(data string) {
	*s = append(*s, data)
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() string {

	index := len(*s) - 1
	val := (*s)[index]
	(*s) = (*s)[:index]
	return val
}

func (s *Stack) Display() {
	fmt.Println(s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
