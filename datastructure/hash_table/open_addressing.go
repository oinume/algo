package hash_table

import "fmt"

const defaultOpenAddressingMaxSize = 53

type bucketState int

const (
	bucketStateEmpty   bucketState = iota
	bucketStateNormal  bucketState = iota
	bucketStateRemoved bucketState = iota
)

type openAddressing[K comparable, V any] struct {
	maxSize int
	size    int
	table   []*bucket[K, V]
}

type bucketKey[K comparable] struct {
	data  K
	state bucketState
}

func newEmptyBucketKey[K comparable]() *bucketKey[K] {
	return &bucketKey[K]{state: bucketStateEmpty}
}

func (k *bucketKey[K]) HashCode() int {
	result := 0
	for _, s := range fmt.Sprint(k.data) {
		result += int(s)
	}
	return result
}

func (k *bucketKey[K]) isEmpty() bool {
	return k.state == bucketStateEmpty
}

func (k *bucketKey[K]) isRemoved() bool {
	return k.state == bucketStateRemoved
}

func (k *bucketKey[K]) setRemoved() {
	k.state = bucketStateRemoved
}

type bucket[K comparable, V any] struct {
	key   *bucketKey[K]
	value V
}

func NewOpenAddressing[K comparable, V any]() Map[K, V] {
	return NewOpenAddressingWithMaxSize[K, V](defaultOpenAddressingMaxSize)
}

func NewOpenAddressingWithMaxSize[K comparable, V any](size int) Map[K, V] {
	table := make([]*bucket[K, V], size)
	for i := 0; i < size; i++ {
		table[i] = &bucket[K, V]{
			key: newEmptyBucketKey[K](),
		}
	}
	return &openAddressing[K, V]{
		maxSize: size,
		table:   table,
	}
}

func (h *openAddressing[K, V]) Put(key K, value V) (V, error) {
	var zero V
	givenKey := &bucketKey[K]{data: key, state: bucketStateNormal}
	index := h.hash(givenKey)
	count := 0
	for k := h.table[index].key; !k.isEmpty() && !k.isRemoved(); k = h.table[index].key {
		if k.data == key {
			old := h.table[index].value
			h.put(givenKey, value, index)
			return old, nil
		}
		if count+1 > h.maxSize {
			return zero, ErrHashTableIsFull
		}
		index = h.rehash(index)
		count++
	}
	h.put(givenKey, value, index)
	h.size++
	return zero, nil
}

func (h *openAddressing[K, V]) put(key *bucketKey[K], value V, index int) {
	h.table[index] = &bucket[K, V]{
		key:   key,
		value: value,
	}
}

func (h *openAddressing[K, V]) Get(key K) (V, error) {
	var zero V
	count := 0
	givenKey := &bucketKey[K]{data: key, state: bucketStateNormal}
	index := h.hash(givenKey)
	for k := h.table[index].key; !k.isEmpty() && !k.isRemoved(); k = h.table[index].key {
		if k.data == key {
			return h.table[index].value, nil
		}
		if count+1 > h.maxSize {
			return zero, ErrNotExists
		}
		index = h.rehash(index)
		count++
	}
	return zero, ErrNotExists
}

func (h *openAddressing[K, V]) Size() int {
	return h.size
}

func (h *openAddressing[K, V]) Remove(key K) (V, error) {
	var zero V
	count := 0
	givenKey := &bucketKey[K]{data: key, state: bucketStateNormal}
	index := h.hash(givenKey)
	for k := h.table[index].key; !k.isEmpty(); k = h.table[index].key {
		if k.data == key {
			k.setRemoved()
			removed := h.table[index].value
			h.table[index].value = zero
			h.size--
			return removed, nil
		}
		if count+1 > h.maxSize {
			return zero, ErrNotExists
		}
		index = h.rehash(index)
		count++
	}
	return zero, ErrNotExists
}

func (h *openAddressing[K, V]) hash(key *bucketKey[K]) int {
	return key.HashCode() % h.maxSize
}

func (h *openAddressing[K, V]) rehash(hash int) int {
	return (hash + 1) % h.maxSize
}
