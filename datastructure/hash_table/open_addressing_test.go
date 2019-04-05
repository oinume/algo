package hash_table

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBucketKey_IsEmpty(t *testing.T) {
	a := assert.New(t)

	empty := &bucketKey{data: emptyKey{}}
	a.True(empty.isEmpty())
}

func TestBucketKey_HashCode(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		key1         interface{}
		key2         interface{}
		sameHashCode bool
	}{
		{key1: 1, key2: 1, sameHashCode: true},
		{key1: "a", key2: "b", sameHashCode: false},
		{key1: "abc", key2: "cba", sameHashCode: true},
	}
	for _, tc := range testCases {
		key1, key2 := &bucketKey{data: tc.key1}, &bucketKey{data: tc.key2}
		a.Equalf(
			tc.sameHashCode,
			key1.HashCode() == key2.HashCode(),
			"sameHashCode=%v", tc.sameHashCode,
		)
	}
}

func TestOpenAddressing_Put(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	hashTable := NewOpenAddressing()

	testCases := []struct {
		key        interface{}
		value      interface{}
		wantReturn interface{}
	}{
		{key: 1, value: 10, wantReturn: nil},
		{key: 2, value: 20, wantReturn: nil},
		{key: 2, value: 30, wantReturn: 20},
	}
	for _, tc := range testCases {
		ret, err := hashTable.Put(tc.key, tc.value)
		r.NoError(err)
		a.EqualValues(tc.wantReturn, ret)
	}
	a.Equal(2, hashTable.Size())
}

func TestOpenAddressing_Put_Rehash(t *testing.T) {
	// TODO: table driven test
	a := assert.New(t)
	r := require.New(t)
	hashTable := NewOpenAddressingWithMaxSize(3)

	ret, err := hashTable.Put("abc", "ABC")
	r.NoError(err)
	a.Nil(ret)
	ret, err = hashTable.Put("cba", "CBA")
	r.NoError(err)
	a.Nil(ret)

	ret, err = hashTable.Get("abc")
	r.NoError(err)
	a.Equal("ABC", ret)

	ret, err = hashTable.Get("cba")
	r.NoError(err)
	a.Equal("CBA", ret)
}

func TestOpenAddressing_Get(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	hashTable := NewOpenAddressing()

	testCases := []struct {
		key   interface{}
		value interface{}
	}{
		{key: 1, value: 10},
		{key: 2, value: 20},
	}
	for _, tc := range testCases {
		_, err := hashTable.Put(tc.key, tc.value)
		r.NoError(err)
		actual, err := hashTable.Get(tc.key)
		r.NoError(err)
		a.Equal(tc.value, actual)
	}
}

func TestOpenAddressing_Remove(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	hashTable := NewOpenAddressing()

	testCases := []struct {
		key    interface{}
		value  interface{}
		remove bool
	}{
		{key: "abc", value: "ABC", remove: true},
		{key: "cba", value: "CBA", remove: false},
	}
	for _, tc := range testCases {
		_, err := hashTable.Put(tc.key, tc.value)
		r.NoError(err)
	}
	size := hashTable.Size()
	r.Equal(len(testCases), size)

	for _, tc := range testCases {
		if tc.remove {
			removed, err := hashTable.Remove(tc.key)
			r.NoError(err)
			a.Equal(tc.value, removed)
			size--
		}
	}
	r.Equal(size, hashTable.Size(), "Size() must be decremented by removal")
}
