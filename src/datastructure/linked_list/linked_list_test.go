package linked_list

import (
	"testing"

	"github.com/oinume/algo/src/datastructure/types"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListAdd(t *testing.T) {
	list := NewLinkedList()
	list.Add(&types.Object{1})
	list.Add(&types.Object{2})
	list.Add(&types.Object{4})

	assert := assert.New(t)
	assert.Equal(3, list.Size())
}

func TestLinkedListInsert(t *testing.T) {
	list := NewLinkedList()
	list.Add(&types.Object{1})
	list.Add(&types.Object{2})
	list.Add(&types.Object{4})
	list.Insert(1, &types.Object{3})

	assert := assert.New(t)
	assert.Equal(4, list.Size())

	expect := NewLinkedList()
	for i := 1; i <= 4; i++ {
		expect.Add(&types.Object{i})
	}
	assert.Equal(expect, list)
}

func TestLinkedListSet(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(&types.Object{i})
	}
	old, err := list.Set(2, &types.Object{300})
	assert.NoError(err)
	assert.Equal(&types.Object{3}, old)
	var value types.Value
	for i := list.Iterator(); i.HasNext(); {
		value, _ = i.Next()
	}
	//t.Logf("list = %v", spew.Sdump(list))
	assert.Equal(&types.Object{300}, value)

	oldFirst, err := list.Set(0, &types.Object{100})
	assert.NoError(err)
	assert.Equal(&types.Object{1}, oldFirst)
	if i := list.Iterator(); i.HasNext() {
		first, _ := i.Next()
		assert.Equal(&types.Object{100}, first)
		second, _ := i.Next()
		assert.Equal(&types.Object{2}, second)
	}
}

func TestLinkedListRemove(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(&types.Object{i})
	}
	assert.True(list.Remove(&types.Object{1}))
	if v, err := list.Iterator().Next(); err == nil {
		assert.Equal(&types.Object{2}, v)
	}

	assert.True(list.Remove(&types.Object{2}))
	assert.True(list.Remove(&types.Object{3}))
	assert.False(list.Remove(&types.Object{100}))
	assert.Equal(0, list.Size())
	assert.False(list.Iterator().HasNext())
	list.Add(&types.Object{10})
}

func TestLinkedListIteratorRemove(t *testing.T) {
	list := NewLinkedList()
	for i := 1; i <= 3; i++ {
		list.Add(&types.Object{i})
	}

	iterator := list.Iterator()
	first, _ := iterator.Next()

	assert := assert.New(t)
	assert.Equal(&types.Object{1}, first)

	for i := 1; i <= 3; i++ {
		removed, err := iterator.Remove()
		assert.NoError(err)
		assert.Equal(&types.Object{i}, removed)
	}
}
