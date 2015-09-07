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

	poped, err := stack.Pop()
	assert.Nil(err)
	assert.Equal(2, stack.Size())
	assert.Equal(Object{Value: 3}, poped)

	stack.Clear()
	assert.Equal(0, stack.Size())
	_, err = stack.Pop()
	assert.True(err != nil)
}
