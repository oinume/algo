package stack

import (
	"fmt"
)

var (
	ErrEmptyStack = fmt.Errorf("stack is empty")
)

type Stack[T any] struct {
	data     []T
	capacity int
}

func New[T any](capacity int) *Stack[T] {
	if capacity <= 0 {
		panic("must be 'capacity' > 0")
	}
	data := make([]T, 0, capacity)
	return &Stack[T]{data: data, capacity: capacity}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		var ret T
		return ret, ErrEmptyStack
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Size() == 0 {
		var ret T
		return ret, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0, s.capacity)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
