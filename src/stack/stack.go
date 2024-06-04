package stack

import "fmt"

type Stack struct {
	size     int // 0 if unlimited
	top      int
	elements []int
}

func NewStack(size int) (*Stack, error) {
	if size < 0 {
		return nil, fmt.Errorf("stack size must be greater than 0")
	}

	return &Stack{
		size:     size,
		top:      -1,
		elements: make([]int, 0, size),
	}, nil
}

func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

func (s *Stack) Push(element int) error {
	if s.size > 0 && s.top == s.size-1 {
		return fmt.Errorf("stack overflow: cannot push more items onto the stack")
	}

	s.top++
	s.elements = append(s.elements, element)

	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return -1, fmt.Errorf("stack underflow: cannot pop from an empty stack")
	}

	element := s.elements[s.top]
	s.elements = s.elements[:s.top]
	s.top--

	return element, nil
}
