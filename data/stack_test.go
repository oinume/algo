package data

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(5)
	stack.Push(Object{Value: 1})
	stack.Push(Object{Value: 2})
}
