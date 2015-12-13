package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListAdd(t *testing.T) {
	list := NewLinkedList()
	list.Add(&Object{1})
	list.Add(&Object{2})
	list.Add(&Object{4})

	assert := assert.New(t)
	assert.Equal(3, list.Size())
}

func TestLinkedListSet(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(&Object{i})
	}
	old, err := list.Set(2, &Object{300})
	assert.NoError(err)
	assert.Equal(&Object{3}, old)
	var value *Object
	for i := list.Iterator(); i.HasNext(); {
		value, _ = i.Next()
	}
	//t.Logf("list = %v", spew.Sdump(list))
	assert.Equal(&Object{300}, value)
}

func TestLinkedListIteratorRemove(t *testing.T) {
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(&Object{i})
	}

	iterator := list.Iterator()
	first, _ := iterator.Next()

	assert := assert.New(t)
	assert.Equal(&Object{1}, first)

	for i := 1; i <= 3; i++ {
		removed, err := iterator.Remove()
		assert.NoError(err)
		assert.Equal(&Object{i}, removed)
	}
}
