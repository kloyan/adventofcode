package common

import "fmt"

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(e interface{}) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	i := len(s.elements) - 1
	last := s.elements[i]
	s.elements = s.elements[:i]

	return last, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
