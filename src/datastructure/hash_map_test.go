package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMapPut(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashMap(10)
	assert.Nil(hashMap.Put(&Object{1}, &Object{1}))
	assert.Equal(1, hashMap.Size())
}

func TestHashMapGet(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashMap(10)
	hashMap.Put(&Object{1}, &Object{1})
	value, err := hashMap.Get(&Object{1})
	assert.NoError(err)
	assert.Equal(&Object{1}, value)
}

func TestHashMapRemove(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashMap(10)
	hashMap.Put(&Object{1}, &Object{1})
	value, err := hashMap.Remove(&Object{1})
	assert.NoError(err)
	assert.Equal(&Object{1}, value)
	assert.Equal(0, hashMap.Size())

	value, err = hashMap.Remove(&Object{100})
	assert.Error(err)
}
