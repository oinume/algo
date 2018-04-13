package datastructure

import (
	"fmt"
	"container/list"
)

const defaultMaxSize = 100

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
			if i := e.Value.(*item); i.key.Get() == key.Get() {
				// Replace an old item with new one
				l.Remove(e)
				l.PushBack(&item{key: key, value: value})
				h.size++
				return i.value
			}
		}
	}
	return nil
}

func (h *hashTableChaining) Get(key Value) (Value, error) {
	index := h.getIndex(key)
	if h.data[index] == nil {
		return nil, fmt.Errorf("not found")
	}
	list := h.data[index]
	for e := list.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item); i.key.Get() == key.Get() {
			return i.value, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (h *hashTableChaining) Remove(key Value) (Value, error) {
	index := h.getIndex(key)
	if h.data[index] == nil {
		return nil, fmt.Errorf("not found")
	}
	list := h.data[index]
	for e := list.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item); i.key.Get() == key.Get() {
			removed := list.Remove(e)
			h.size--
			return removed.(*item).value, nil
		}
	}
	return nil, fmt.Errorf("not found")
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
