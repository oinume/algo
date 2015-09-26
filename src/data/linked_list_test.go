package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListAdd(t *testing.T) {
	list := NewLinkedList()
	list.Add(Object{1})
	list.Add(Object{2})
	list.Add(Object{4})

	assert := assert.New(t)
	assert.Equal(3, list.Size())
}

func TestLinkedListIteratorRemove(t *testing.T) {
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(Object{i})
	}

	iterator := list.Iterator()
	first, _ := iterator.Next()

	assert := assert.New(t)
	assert.Equal(Object{1}, first)

	for i := 1; i <= 3; i++ {
		removed, err := iterator.Remove()
		assert.NoError(err)
		assert.Equal(Object{i}, removed)
	}
}
