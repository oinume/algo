package hash_table

import (
	"fmt"
	"reflect"

	"github.com/oinume/algo/src/datastructure/types"
)

const defaultOpenAddressingMaxSize = 53

type (
	emptyKey   struct{}
	removedKey struct{}
)

type openAddressing struct {
	maxSize int
	size    int
	table   []*bucket
}

type bucketKey struct {
	data types.Value
}

func (k *bucketKey) HashCode() int {
	result := 0
	for _, s := range fmt.Sprint(k.data) {
		result += int(s)
	}
	return result
}

func (k *bucketKey) isEmpty() bool {
	if k.data == nil {
		return true
	}
	if _, ok := k.data.Get().(emptyKey); ok {
		return true
	}
	return false
}

func (k *bucketKey) isRemoved() bool {
	if _, ok := k.data.Get().(removedKey); ok {
		return true
	}
	return false
}

func (k *bucketKey) setRemoved() {
	k.data = &types.Object{Value: removedKey{}}
}

type bucket struct {
	key   *bucketKey
	value types.Value
}

func NewOpenAddressing() types.Map {
	return NewOpenAddressingWithMaxSize(defaultOpenAddressingMaxSize)
}

func NewOpenAddressingWithMaxSize(size int) types.Map {
	table := make([]*bucket, size)
	for i := 0; i < size; i++ {
		table[i] = &bucket{
			key:   &bucketKey{data: &types.Object{Value: emptyKey{}}},
			value: nil,
		}
	}
	hashTable := &openAddressing{
		maxSize: size,
		table:   table,
	}
	return hashTable
}

func (h *openAddressing) Put(key types.Value, value types.Value) (types.Value, error) {
	if key == nil {
		return nil, ErrKeyMustNotBeNil
	}

	givenKey := &bucketKey{data: key}
	index := h.hash(givenKey)
	count := 0
	for k := h.table[index].key; !k.isEmpty() && !k.isRemoved(); k = h.table[index].key {
		if reflect.DeepEqual(givenKey.data, k.data) {
			// Already exists, replace it with a new value
			old := h.table[index].value
			h.put(givenKey, value, index)
			return old, nil
		}
		if count+1 > h.maxSize {
			return nil, ErrHashTableIsFull
		}
		index = h.rehash(index)
		count++
	}
	h.put(givenKey, value, index)
	h.size++
	return nil, nil
}

func (h *openAddressing) put(key *bucketKey, value types.Value, index int) {
	h.table[index] = &bucket{
		key:   key,
		value: value,
	}
}

func (h *openAddressing) Get(key types.Value) (types.Value, error) {
	count := 0
	givenKey := &bucketKey{data: key}
	index := h.hash(givenKey)
	// わかりにくいので for i := 0; i < h.maxSize; i++ {} にする
	for k := h.table[index].key; !k.isEmpty() && !k.isRemoved(); k = h.table[index].key {
		if reflect.DeepEqual(givenKey.data, k.data) {
			// Found
			return h.table[index].value, nil
		}
		if count+1 > h.maxSize {
			return nil, ErrNotExists
		}
		index = h.rehash(index)
		count++
	}
	return nil, ErrNotExists
}

func (h *openAddressing) Size() int {
	return h.size
}

func (h *openAddressing) Remove(key types.Value) (types.Value, error) {
	count := 0
	givenKey := &bucketKey{data: key}
	index := h.hash(givenKey)
	for k := h.table[index].key; !k.isEmpty(); k = h.table[index].key {
		if reflect.DeepEqual(givenKey.data, k.data) {
			// Found
			k.setRemoved()
			removed := h.table[index].value
			h.table[index].value = nil
			h.size--
			return removed, nil
		}
		if count+1 > h.maxSize {
			return nil, ErrNotExists
		}
		index = h.rehash(index)
		count++
	}
	return nil, ErrNotExists
}

func (h *openAddressing) hash(key *bucketKey) int {
	return key.HashCode() % h.maxSize
}

func (h *openAddressing) rehash(hash int) int {
	return (hash + 1) % h.maxSize
}
