package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMapPut(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashMap(10)
	assert.Nil(hashMap.Put(&Object{1}, &Object{1}))
}

func TestHashMapGet(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashMap(10)
	hashMap.Put(&Object{1}, &Object{1})
	value, err := hashMap.Get(&Object{1})
	assert.Nil(err)
	assert.Equal(&Object{1}, value)
}
