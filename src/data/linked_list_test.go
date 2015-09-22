package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()
	list.Add(Object{1})
	list.Add(Object{2})
	list.Add(Object{4})

	assert := assert.New(t)
	assert.Equal(3, list.Size())

	iterator := list.Iterator()
	iterator.Next()
	iterator.Next()
	last, _ := iterator.Next()
	assert.Equal(Object{4}, last)
	assert.Equal(false, iterator.HasNext())
}
