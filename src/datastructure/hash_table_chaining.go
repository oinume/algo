package datastructure

import (
	"fmt"
	"container/list"
	"reflect"
	"errors"
)

const defaultMaxSize = 100
var ErrKeyNotExists = errors.New("key not exists")

type hashTableChaining struct {
	maxSize int
	size int
	data []*list.List
}

type item struct {
	key Value
	value Value
}

func NewHashTableChaining(maxSize int) Map {
	if maxSize <= 0 {
		maxSize = defaultMaxSize
	}
	return &hashTableChaining{
		maxSize: maxSize,
		size: 0,
		data: make([]*list.List, maxSize),
	}
}

func (h *hashTableChaining) Put(key Value, value Value) Value {
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
			if i := e.Value.(*item); reflect.DeepEqual(i.key.Get(), key.Get()) {
				// Replace an old item with new one
				l.Remove(e)
				l.PushBack(&item{key: key, value: value})
				return i.value
			}
		}
		l.PushBack(&item{key: key, value: value})
		h.size++
	}
	return nil
}

func (h *hashTableChaining) Get(key Value) (Value, error) {
	index := h.getIndex(key)
	if h.data[index] == nil {
		return nil, ErrKeyNotExists
	}
	list := h.data[index]
	for e := list.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item); i.key.Get() == key.Get() {
			return i.value, nil
		}
	}
	return nil, ErrKeyNotExists
}

func (h *hashTableChaining) Remove(key Value) (Value, error) {
	index := h.getIndex(key)
	if h.data[index] == nil {
		return nil, ErrKeyNotExists
	}
	list := h.data[index]
	for e := list.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item); i.key.Get() == key.Get() {
			removed := list.Remove(e)
			h.size--
			return removed.(*item).value, nil
		}
	}
	return nil, ErrKeyNotExists
}

func (h *hashTableChaining) Size() int {
	return h.size
}

func (h *hashTableChaining) calculateHashCode(v Value) int {
	result := 0
	for _, s := range fmt.Sprint(v) {
		result += int(s)
	}
	return result
}

func (h *hashTableChaining) getIndex(key Value) int {
	k, ok := key.(Hashable)
	var hashCode int
	if ok {
		hashCode = k.HashCode()
	} else {
		hashCode = h.calculateHashCode(key)
	}
	return hashCode % h.maxSize
}

// TODO: Replace getIndex by this
func (h *hashTableChaining) find(key Value) (*list.List, error) {
	k, ok := key.(Hashable)
	var hashCode int
	if ok {
		hashCode = k.HashCode()
	} else {
		hashCode = h.calculateHashCode(key)
	}
	index := hashCode % h.maxSize
	l := h.data[index]
	if l == nil {
		return nil, ErrKeyNotExists
	}
	return l, nil
}
