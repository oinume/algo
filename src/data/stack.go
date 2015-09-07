package data

import (
	"fmt"
)

type Stack struct {
	data     []Object
	capacity int
}

func NewStack(capacity int) *Stack {
	data := make([]Object, 0, capacity)
	return &Stack{data: data, capacity: capacity}
}

func (s *Stack) Push(o Object) {
	s.data = append(s.data, o)
}

func (s *Stack) Pop() (Object, error) {
	if s.Size() == 0 {
		return Object{}, fmt.Errorf("Stack is empty.")
	}
	o := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return o, nil
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = make([]Object, 0, s.capacity)
}
