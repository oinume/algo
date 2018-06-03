package hash_table

import (
	"testing"

	"github.com/oinume/algo/src/datastructure/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBucketKey_IsEmpty(t *testing.T) {
	a := assert.New(t)

	empty := &bucketKey{data: &types.Object{Value: emptyKey{}}}
	a.True(empty.isEmpty())
}

func TestBucketKey_HashCode(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		key1         types.Value
		key2         types.Value
		sameHashCode bool
	}{
		{key1: &types.Object{1}, key2: &types.Object{1}, sameHashCode: true},
		{key1: &types.Object{"a"}, key2: &types.Object{"b"}, sameHashCode: false},
		{key1: &types.Object{"abc"}, key2: &types.Object{"cba"}, sameHashCode: true},
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
		key        types.Value
		value      types.Value
		wantReturn types.Value
	}{
		{key: &types.Object{1}, value: &types.Object{10}, wantReturn: nil},
		{key: &types.Object{2}, value: &types.Object{20}, wantReturn: nil},
		{key: &types.Object{2}, value: &types.Object{30}, wantReturn: &types.Object{20}},
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

	ret, err := hashTable.Put(&types.Object{"abc"}, &types.Object{"ABC"})
	r.NoError(err)
	a.Nil(ret)
	ret, err = hashTable.Put(&types.Object{"cba"}, &types.Object{"CBA"})
	r.NoError(err)
	a.Nil(ret)

	ret, err = hashTable.Get(&types.Object{"abc"})
	r.NoError(err)
	a.Equal(&types.Object{"ABC"}, ret)

	ret, err = hashTable.Get(&types.Object{"cba"})
	r.NoError(err)
	a.Equal(&types.Object{"CBA"}, ret)
}

func TestOpenAddressing_Get(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	hashTable := NewOpenAddressing()

	testCases := []struct {
		key   types.Value
		value types.Value
	}{
		{key: &types.Object{1}, value: &types.Object{10}},
		{key: &types.Object{2}, value: &types.Object{20}},
	}
	for _, tc := range testCases {
		_, err := hashTable.Put(tc.key, tc.value)
		r.NoError(err)
		actual, err := hashTable.Get(tc.key)
		a.Equal(tc.value, actual)
	}
}
