package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	assert := assert.New(t)

	stack := createStack(5)
	assert.Equal(3, stack.Size())
}

func TestStackPop(t *testing.T) {
	assert := assert.New(t)
	stack := createStack(5)

	poped, err := stack.Pop()
	assert.Nil(err)
	assert.Equal(2, stack.Size())
	assert.Equal(3, poped)

	stack.Clear()
	assert.Equal(0, stack.Size())
	_, err = stack.Pop()
	assert.NotNil(err)
}

func createStack(capacity int) *Stack {
	stack := NewStack(capacity)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	return stack
}
