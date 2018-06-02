package hash_table

import (
	"testing"

	"github.com/oinume/algo/src/datastructure/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAddressing_Put(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	hashTable := NewOpenAddressing()

	testCases := []struct {
		key        types.Value
		value      types.Value
		wantReturn types.Value
	}{
		{
			key: &types.Object{1}, value: &types.Object{10}, wantReturn: nil,
		},
		{
			key: &types.Object{2}, value: &types.Object{20}, wantReturn: nil,
		},
	}
	for _, tc := range testCases {
		ret, err := hashTable.Put(tc.key, tc.value)
		r.NoError(err)
		a.EqualValues(tc.wantReturn, ret)
	}
	a.Equal(len(testCases), hashTable.Size())
}

func TestOpenAddressing_Get(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	hashTable := NewOpenAddressing()
	hashTable.Put(&types.Object{1}, &types.Object{10})
	hashTable.Put(&types.Object{2}, &types.Object{20})
	a.Equal(2, hashTable.Size())

	actual, err := hashTable.Get(&types.Object{2})
	r.NoError(err)
	a.Equal(&types.Object{20}, actual)
}
