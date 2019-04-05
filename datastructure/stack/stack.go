package stack

import (
	"fmt"
)

var (
	ErrEmptyStack = fmt.Errorf("stack is empty")
)

type Stack struct {
	data     []interface{}
	capacity int
}

func NewStack(capacity int) *Stack {
	if capacity <= 0 {
		panic("must be 'capacity' > 0")
	}
	data := make([]interface{}, 0, capacity)
	return &Stack{data: data, capacity: capacity}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Size() == 0 {
		return nil, ErrEmptyStack
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.Size() == 0 {
		return nil, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = make([]interface{}, 0, s.capacity)
}
