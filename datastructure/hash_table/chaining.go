package hash_table

import (
	"container/list"
	"fmt"

	"github.com/oinume/algo/datastructure/types"
)

const defaultMaxSize = 100

type chaining[K comparable, V any] struct {
	maxSize int
	size    int
	data    []*list.List
}

type item[K comparable, V any] struct {
	key   K
	value V
}

func NewChaining[K comparable, V any](maxSize int) types.Map[K, V] {
	if maxSize <= 0 {
		maxSize = defaultMaxSize
	}
	return &chaining[K, V]{
		maxSize: maxSize,
		size:    0,
		data:    make([]*list.List, maxSize),
	}
}

func (h *chaining[K, V]) Put(key K, value V) (V, error) {
	var zero V
	index := h.getIndex(key)
	if h.data[index] == nil {
		l := list.New()
		l.PushBack(&item[K, V]{key: key, value: value})
		h.data[index] = l
		h.size++
	} else {
		l := h.data[index]
		for e := l.Front(); e != nil; e = e.Next() {
			if i := e.Value.(*item[K, V]); i.key == key {
				old := i.value
				l.Remove(e)
				l.PushBack(&item[K, V]{key: key, value: value})
				return old, nil
			}
		}
		l.PushBack(&item[K, V]{key: key, value: value})
		h.size++
	}
	return zero, nil
}

func (h *chaining[K, V]) Get(key K) (V, error) {
	var zero V
	index := h.getIndex(key)
	if h.data[index] == nil {
		return zero, ErrNotExists
	}
	l := h.data[index]
	for e := l.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item[K, V]); i.key == key {
			return i.value, nil
		}
	}
	return zero, ErrNotExists
}

func (h *chaining[K, V]) Remove(key K) (V, error) {
	var zero V
	index := h.getIndex(key)
	if h.data[index] == nil {
		return zero, ErrNotExists
	}
	l := h.data[index]
	for e := l.Front(); e != nil; e = e.Next() {
		if i := e.Value.(*item[K, V]); i.key == key {
			l.Remove(e)
			h.size--
			return i.value, nil
		}
	}
	return zero, ErrNotExists
}

func (h *chaining[K, V]) Size() int {
	return h.size
}

func (h *chaining[K, V]) calculateHashCode(key K) int {
	result := 0
	for _, s := range fmt.Sprint(key) {
		result += int(s)
	}
	return result
}

func (h *chaining[K, V]) getIndex(key K) int {
	return h.calculateHashCode(key) % h.maxSize
}
