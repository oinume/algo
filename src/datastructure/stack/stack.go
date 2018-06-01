package stack

import (
	"fmt"

	"github.com/oinume/algo/src/datastructure/types"
)

type Stack struct {
	data     []types.Value
	capacity int
}

func NewStack(capacity int) *Stack {
	if capacity <= 0 {
		panic("Must be 'capacity' > 0")
	}
	data := make([]types.Value, 0, capacity)
	return &Stack{data: data, capacity: capacity}
}

func (s *Stack) Push(v types.Value) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (types.Value, error) {
	if s.Size() == 0 {
		return nil, fmt.Errorf("Stack is empty.")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *Stack) Peek() (types.Value, error) {
	if s.Size() == 0 {
		return nil, fmt.Errorf("Stack is empty.")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = make([]types.Value, 0, s.capacity)
}
