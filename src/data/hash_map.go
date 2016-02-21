package data

import (
	"fmt"
)

const defaultMaxSize = 100

type hashMap struct {
	maxSize int
	data []Value
}

func NewHashMap(maxSize int) Map {
	if maxSize <= 0 {
		maxSize = defaultMaxSize
	}
	return &hashMap{
		maxSize: maxSize,
		data: make([]Value, maxSize),
	}
}

func (h *hashMap) Put(key Value, value Value) Value {
	k, ok := key.(Hashable)
	var hashCode int
	if ok {
		hashCode = k.HashCode()
	} else {
		hashCode = h.hashCode(key)
	}
	loc := hashCode % h.maxSize
	h.data[loc] = value
	return value
}

func (h *hashMap) Get(key Value) (Value, error) {
	return nil, nil
}

func (h *hashMap) hashCode(v Value) int {
	result := 0
	for _, s := range fmt.Sprint(v) {
		result += int(s)
	}
	return result
}
