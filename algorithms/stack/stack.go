package stack

import "errors"

type Stack struct {
	data []int
}

func (s *Stack) Push(value ...int) bool {
	s.data = append(s.data, value...)
	return true
}

func (s *Stack) Pop() (int, error) {
	length := len(s.data)

	if length <= 0 {
		return 0, errors.New("The stack is empty")
	}

	last := s.data[length-1]

	s.data = s.data[0 : length-1]

	return last, nil
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Top() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("The stack is empty.")
	}
	return s.data[len(s.data)-1], nil
}

func IsValid(s string) bool {

	stack := &Stack{}

	for _, char := range s {

		if stack.Empty() {
			stack.Push(int(char))
		} else {
			top, _ := stack.Top()

			if (top == '(' && char == ')') ||
				(top == '[' && char == ']') ||
				(top == '{' && char == '}') {
				_, _ = stack.Pop()
			} else {
				stack.Push(int(char))
			}
		}
	}
	return stack.Empty()
}
