package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTableChaining_Put(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashTableChaining(10)
	assert.Nil(hashMap.Put(&Object{1}, &Object{1}))
	assert.Equal(1, hashMap.Size())
}

func TestHashTableChaining_Put_Collision(t *testing.T) {
	// TODO: test same hash code
	assert := assert.New(t)
	hashMap := NewHashTableChaining(10)
	assert.Nil(hashMap.Put(&Object{1}, &Object{1}))
	assert.Equal(1, hashMap.Size())
}

func TestHashTableChaining_Get(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashTableChaining(10)
	hashMap.Put(&Object{1}, &Object{1})
	value, err := hashMap.Get(&Object{1})
	assert.NoError(err)
	assert.Equal(&Object{1}, value)
}

func TestHashTableChaining_Remove(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashTableChaining(10)
	hashMap.Put(&Object{1}, &Object{1})
	value, err := hashMap.Remove(&Object{1})
	assert.NoError(err)
	assert.Equal(&Object{1}, value)
	assert.Equal(0, hashMap.Size())

	value, err = hashMap.Remove(&Object{100})
	assert.Error(err)
}
