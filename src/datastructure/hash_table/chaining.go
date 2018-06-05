package hash_table

import (
	"container/list"
	"fmt"
	"reflect"

	"github.com/oinume/algo/src/datastructure/types"
)

const defaultMaxSize = 100

type chaining struct {
	maxSize int
	size    int
	data    []*list.List
}

type item struct {
	key   interface{}
	value interface{}
}

func NewChaining(maxSize int) types.Map {
	if maxSize <= 0 {
		maxSize = defaultMaxSize
	}
	return &chaining{
		maxSize: maxSize,
		size:    0,
		data:    make([]*list.List, maxSize),
	}
}

func (h *chaining) Put(key interface{}, value interface{}) (interface{}, error) {
	if key == nil {
		return nil, ErrKeyMustNotBeNil
	}
	index := h.getIndex(key)
	if h.data[index] == nil {
		// Put as new
		l := list.New()
		l.PushBack(&item{key: key, value: value})
		h.data[index] = l
		h.size++
	} else {
		l := h.data[index]
		for e := l.Front(); e != nil; e = e.Next() {
			if i := e.Value.(*item); reflect.DeepEqual(i.key, key) {
				// Replace an old item with new one
				l.Remove(e)
				l.PushBack(&item{key: key, value: value})
				return i.value, nil
			}
		}
		l.PushBack(&item{key: key, value: value})
		h.size++
	}
	return nil, nil
}

func (h *chaining) Get(key interface{}) (interface{}, error) {
	index := h.getIndex(key)
	if h.data[index] == nil {
		return nil, ErrNotExists
	}
	list := h.data[index]
	for e := list.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item); i.key == key {
			return i.value, nil
		}
	}
	return nil, ErrNotExists
}

func (h *chaining) Remove(key interface{}) (interface{}, error) {
	index := h.getIndex(key)
	if h.data[index] == nil {
		return nil, ErrNotExists
	}
	list := h.data[index]
	for e := list.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item); i.key == key {
			removed := list.Remove(e)
			h.size--
			return removed.(*item).value, nil
		}
	}
	return nil, ErrNotExists
}

func (h *chaining) Size() int {
	return h.size
}

func (h *chaining) calculateHashCode(v interface{}) int {
	result := 0
	for _, s := range fmt.Sprint(v) {
		result += int(s)
	}
	return result
}

func (h *chaining) getIndex(key interface{}) int {
	k, ok := key.(types.Hashable)
	var hashCode int
	if ok {
		hashCode = k.HashCode()
	} else {
		hashCode = h.calculateHashCode(key)
	}
	return hashCode % h.maxSize
}

// TODO: Replace getIndex by this
func (h *chaining) find(key interface{}) (*list.List, error) {
	k, ok := key.(types.Hashable)
	var hashCode int
	if ok {
		hashCode = k.HashCode()
	} else {
		hashCode = h.calculateHashCode(key)
	}
	index := hashCode % h.maxSize
	l := h.data[index]
	if l == nil {
		return nil, ErrNotExists
	}
	return l, nil
}
