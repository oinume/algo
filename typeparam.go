// run -gcflags=-G=3
package main

import "fmt"

var (
	ErrEmptyStack = fmt.Errorf("stack is empty")
)

type stack[T any] struct {
	data     []T
	capacity int
}

func newStack[T any](capacity int) *stack[T] {
	if capacity <= 0 {
		panic("must be 'capacity' > 0")
	}
	data := make([]T, 0, capacity)
	return &stack[T]{data: data, capacity: capacity}
}

func (s *stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *stack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		var zero T
		return zero, ErrEmptyStack
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *stack[T]) Peek() (T, error) {
	if s.Size() == 0 {
		var zero T
		return zero, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

func (s *stack[T]) Size() int {
	return len(s.data)
}

func main() {
	s := newStack[int](10)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("s = %+v\n", s)

	_, err := s.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("s = %+v\n", s)
}
