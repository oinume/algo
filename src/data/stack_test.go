package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)

	stack := NewStack(5)
	stack.Push(Object{Value: 1})
	stack.Push(Object{Value: 2})
	stack.Push(Object{Value: 3})
	assert.Equal(3, stack.Size())

	poped := stack.Pop()
	assert.Equal(2, stack.Size())
	assert.Equal(Object{Value: 3}, poped)
}
