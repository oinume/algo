package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListAdd(t *testing.T) {
	list := NewLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(4)

	assert := assert.New(t)
	assert.Equal(3, list.Size())
}

func TestLinkedListInsert(t *testing.T) {
	list := NewLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(4)
	list.Insert(1, 3)

	assert := assert.New(t)
	assert.Equal(4, list.Size())

	expect := NewLinkedList()
	for i := 1; i <= 4; i++ {
		expect.Add(i)
	}
	assert.Equal(expect, list)
}

func TestLinkedListSet(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(i)
	}
	old, err := list.Set(2, 300)
	assert.NoError(err)
	assert.Equal(3, old)
	var value interface{}
	for i := list.Iterator(); i.HasNext(); {
		value, _ = i.Next()
	}
	//t.Logf("list = %v", spew.Sdump(list))
	assert.Equal(300, value)

	oldFirst, err := list.Set(0, 100)
	assert.NoError(err)
	assert.Equal(1, oldFirst)
	if i := list.Iterator(); i.HasNext() {
		first, _ := i.Next()
		assert.Equal(100, first)
		second, _ := i.Next()
		assert.Equal(2, second)
	}
}

func TestLinkedListRemove(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(i)
	}
	assert.True(list.Remove(1))
	if v, err := list.Iterator().Next(); err == nil {
		assert.Equal(2, v)
	}

	assert.True(list.Remove(2))
	assert.True(list.Remove(3))
	assert.False(list.Remove(100))
	assert.Equal(0, list.Size())
	assert.False(list.Iterator().HasNext())
	list.Add(10)
}

func TestLinkedListIteratorRemove(t *testing.T) {
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(i)
	}

	iterator := list.Iterator()
	first, _ := iterator.Next()

	assert := assert.New(t)
	assert.Equal(1, first)

	for i := 1; i <= 3; i++ {
		removed, err := iterator.Remove()
		assert.NoError(err)
		assert.Equal(i, removed)
	}
}
