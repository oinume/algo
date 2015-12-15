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

func TestLinkedListInsert(t *testing.T) {
	list := NewLinkedList()
	list.Add(&Object{1})
	list.Add(&Object{2})
	list.Add(&Object{4})
	list.Insert(1, &Object{3})

	assert := assert.New(t)
	assert.Equal(4, list.Size())

	expect := NewLinkedList()
	for i := 1; i <= 4; i++ {
		expect.Add(&Object{i})
	}
	assert.Equal(expect, list)
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
	var value Value
	for i := list.Iterator(); i.HasNext(); {
		value, _ = i.Next()
	}
	//t.Logf("list = %v", spew.Sdump(list))
	assert.Equal(&Object{300}, value)
}

func TestLinkedListRemove(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(&Object{i})
	}
	assert.True(list.Remove(&Object{1}))
	assert.True(list.Remove(&Object{2}))
	assert.True(list.Remove(&Object{3}))
	assert.False(list.Remove(&Object{100}))
	assert.Equal(0, list.Size())
	assert.False(list.Iterator().HasNext())
	list.Add(&Object{10})
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
