package stack

func (s *Stack) BalancedParenthesis(input string) bool {
	for i := 0; i < len(input); i++ {
		ch := string(input[i])

		if s.IsEmpty() {
			s.Push(ch)
		} else if s.Peek() == "(" && ch == ")" || s.Peek() == "{" && ch == "}" || s.Peek() == "[" && ch == "]" {
			s.Pop()
		} else {
			s.Push(ch)
		}
	}
	return s.IsEmpty()
}

func (s *Stack) ReverseString(input string) string {
	result := ""

	for _, ch := range input {
		s.Push(string(ch))
	}

	for !s.IsEmpty() {
		result += s.Pop()
	}
	return result
}
