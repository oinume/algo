package hash_table

import "github.com/oinume/algo/src/datastructure/types"

const defaultOpenAddressingMaxSize = 53

type hashTableOpenAddressing struct {
	size         int
	elementCount int
	table        []*bucket
}

type bucketKey struct {
	data types.Value
}

func (k *bucketKey) HashCode() int {
	panic("implement")
}

type bucket struct {
	key   *bucketKey
	value types.Value
}

func NewOpenAddresssing() *hashTableOpenAddressing {
	return NewOpenAddressingWithMaxSize(defaultOpenAddressingMaxSize)
}

func NewOpenAddressingWithMaxSize(size int) *hashTableOpenAddressing {
	table := make([]*bucket, size)
	for i := 0; i < size; i++ {
		table[i] = &bucket{
			key:   &bucketKey{data: nil},
			value: nil,
		}
	}
	hashTable := &hashTableOpenAddressing{
		size:  size,
		table: table,
	}
	return hashTable
}

func (hashTableOpenAddressing) Put(key types.Value, value types.Value) types.Value {
	panic("implement me")
}

func (hashTableOpenAddressing) Get(key types.Value) (types.Value, error) {
	panic("implement me")
}

func (hashTableOpenAddressing) Size() int {
	panic("implement me")
}

func (hashTableOpenAddressing) Remove(key types.Value) (types.Value, error) {
	panic("implement me")
}

func (h *hashTableOpenAddressing) hash(key *bucketKey) int {
	return key.HashCode() % h.size
}

func (h *hashTableOpenAddressing) rehash(hash int) int {
	return (hash + 1) % h.size
}
