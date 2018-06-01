package hash_table

import (
	"testing"

	"github.com/oinume/algo/src/datastructure/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type hashable struct {
	object types.Object
}

func (h *hashable) Get() interface{} {
	return h.object
}

func (h *hashable) Receive(v interface{}) error {
	return h.object.Receive(v)
}

func (h *hashable) String() string {
	return h.object.String()
}

func (h *hashable) Int() int {
	return h.object.Int()
}

// Always return same hash code
func (h *hashable) HashCode() int {
	return 1
}

func TestHashTableChaining_Put(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewChaining(10)
	assert.Nil(hashMap.Put(&types.Object{1}, &types.Object{1}))
	assert.Equal(1, hashMap.Size())
}

func TestHashTableChaining_Put_Collision(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	table := NewChaining(10)
	table.Put(&hashable{types.Object{1}}, &hashable{types.Object{10}})
	table.Put(&hashable{types.Object{2}}, &hashable{types.Object{20}})
	a.Equal(2, table.Size())

	actual, err := table.Get(&hashable{types.Object{2}})
	r.NoError(err)
	a.Equal(&hashable{types.Object{20}}, actual)
}

func TestHashTableChaining_Put_Collision_Exists(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	table := NewChaining(10)
	table.Put(&hashable{types.Object{1}}, &hashable{types.Object{10}})
	table.Put(&hashable{types.Object{1}}, &hashable{types.Object{11}})
	table.Put(&hashable{types.Object{2}}, &hashable{types.Object{20}})
	a.Equal(2, table.Size())

	actual, err := table.Get(&hashable{types.Object{1}})
	r.NoError(err)
	a.Equal(&hashable{types.Object{11}}, actual)
}

func TestHashTableChaining_Get(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewChaining(10)
	hashMap.Put(&types.Object{1}, &types.Object{1})
	value, err := hashMap.Get(&types.Object{1})
	assert.NoError(err)
	assert.Equal(&types.Object{1}, value)
}

func TestHashTableChaining_Remove(t *testing.T) {
	assert := assert.New(t)
	hashMap := NewChaining(10)
	hashMap.Put(&types.Object{1}, &types.Object{1})
	value, err := hashMap.Remove(&types.Object{1})
	assert.NoError(err)
	assert.Equal(&types.Object{1}, value)
	assert.Equal(0, hashMap.Size())

	value, err = hashMap.Remove(&types.Object{100})
	assert.Error(err)
}
