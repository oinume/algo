package hash_table

import (
	"reflect"

	"fmt"

	"github.com/oinume/algo/src/datastructure/types"
)

const defaultOpenAddressingMaxSize = 53

type (
	emptyKey   struct{} // TODO: Use this
	deletedKey struct{}
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
	return k.data == nil
}

func (k *bucketKey) isDeleted() bool {
	return false // TODO: implement
}

type bucket struct {
	key   *bucketKey
	value types.Value
}

func NewOpenAddresssing() *openAddressing {
	return NewOpenAddressingWithMaxSize(defaultOpenAddressingMaxSize)
}

func NewOpenAddressingWithMaxSize(size int) *openAddressing {
	table := make([]*bucket, size)
	for i := 0; i < size; i++ {
		table[i] = &bucket{
			key:   &bucketKey{data: nil},
			value: nil,
		}
	}
	hashTable := &openAddressing{
		maxSize: size,
		table:   table,
	}
	return hashTable
}

func (h *openAddressing) Put(key types.Value, value types.Value) types.Value {
	if key == nil {
		panic("key cannot be nil") // TODO: return error?
	}

	givenKey := &bucketKey{data: key}
	hashCode := givenKey.HashCode()
	count := 0
	for k := h.table[hashCode].key; !k.isEmpty() && !k.isDeleted(); {
		if reflect.DeepEqual(givenKey.data, k.data) {
			// Already exists, replace it with a new value
			h.put(givenKey, value, hashCode)
			return h.table[hashCode].value
		}
		if count+1 > h.maxSize {
			panic("HashTable is full.") // TODO: return error?
		}
		hashCode = h.rehash(hashCode)
		count++
	}
	h.put(givenKey, value, hashCode)
	h.size++
	return nil
}

func (h *openAddressing) put(key *bucketKey, value types.Value, hashCode int) {
	h.table[hashCode] = &bucket{
		key:   key,
		value: value,
	}
}

func (h *openAddressing) Get(key types.Value) (types.Value, error) {
	count := 0
	givenKey := &bucketKey{data: key}
	hashCode := givenKey.HashCode()
	// わかりにくいので for i := 0; i < h.maxSize; i++ {} にする
	for k := h.table[hashCode].key; !k.isEmpty() && !k.isDeleted(); {
		if reflect.DeepEqual(givenKey.data, k.data) {
			// Found
			return h.table[hashCode].value, nil
		}
		if count+1 > h.maxSize {
			return nil, ErrKeyNotExists
		}
		hashCode = h.rehash(hashCode)
		count++
	}
	return nil, ErrKeyNotExists
}

func (h *openAddressing) Size() int {
	return h.size
}

func (openAddressing) Remove(key types.Value) (types.Value, error) {
	panic("implement me")
}

func (h *openAddressing) hash(key *bucketKey) int {
	return key.HashCode() % h.maxSize
}

func (h *openAddressing) rehash(hash int) int {
	return (hash + 1) % h.maxSize
}
