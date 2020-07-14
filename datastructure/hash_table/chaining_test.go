package hash_table_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oinume/algo/datastructure/hash_table"
)

//type hashable struct {
//	object types.Object
//}
//
//func (h *hashable) Get() interface{} {
//	return h.object
//}
//
//func (h *hashable) Receive(v interface{}) error {
//	return h.object.Receive(v)
//}
//
//func (h *hashable) String() string {
//	return h.object.String()
//}
//
//func (h *hashable) Int() int {
//	return h.object.Int()
//}

//// Always return same hash code
//func (h *hashable) HashCode() int {
//	return 1
//}

func TestHashTableChaining_Put(t *testing.T) {
	h := hash_table.NewChaining(10)
	_, err := h.Put(1, 1)
	if err != nil {
		t.Fatalf("Put: unexpected error: %v", err)
	}
	if got, want := 1, h.Size(); got != want {
		t.Fatalf("unexpected Size: got=%v, want=%v", got, want)
	}
}

func TestHashTableChaining_Put_Collision(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	table := NewChaining(10)
	_, _ = table.Put("abc", "ABC")
	_, _ = table.Put("cba", "CBA")
	a.Equal(2, table.Size())

	actual, err := table.Get("cba")
	r.NoError(err)
	a.Equal("CBA", actual)
}

func TestHashTableChaining_Put_Collision_Exists(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	table := NewChaining(10)
	_, _ = table.Put("abc", "ABC")
	_, _ = table.Put("abc", "AABBCC")
	_, _ = table.Put("cba", "CBA")
	a.Equal(2, table.Size())

	actual, err := table.Get("abc")
	r.NoError(err)
	a.Equal("AABBCC", actual)
}

func TestHashTableChaining_Get(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	hashMap := NewChaining(10)
	_, err := hashMap.Put(1, 1)
	r.NoError(err)
	value, err := hashMap.Get(1)
	a.NoError(err)
	a.Equal(1, value)
}

func TestHashTableChaining_Remove(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	hashMap := NewChaining(10)

	_, err := hashMap.Put(1, 1)
	r.NoError(err)
	value, err := hashMap.Remove(1)
	r.NoError(err)
	a.Equal(1, value)
	a.Equal(0, hashMap.Size())

	_, err = hashMap.Remove(100)
	a.Error(err)
}
