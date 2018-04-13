package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type hashableObject struct {
	object Object
}

func (ho *hashableObject) Get() interface{} {
	return ho.object
}

func (ho *hashableObject) Receive(v interface{}) error {
	return ho.object.Receive(v)
}

func (ho *hashableObject) String() string {
	return ho.object.String()
}

func (ho *hashableObject) Int() int {
	return ho.object.Int()
}

func (ho *hashableObject) HashCode() int {
	return 1
}

func TestHashTableChaining_Put(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewHashTableChaining(10)
	assert.Nil(hashMap.Put(&Object{1}, &Object{1}))
	assert.Equal(1, hashMap.Size())
}

func TestHashTableChaining_Put_Collision(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	table := NewHashTableChaining(10)
	table.Put(&hashableObject{Object{1}}, &hashableObject{Object{10}})
	table.Put(&hashableObject{Object{2}}, &hashableObject{Object{20}})
	a.Equal(2, table.Size())

	actual, err := table.Get(&hashableObject{Object{2}})
	r.NoError(err)
	a.Equal(&hashableObject{Object{20}}, actual)
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
