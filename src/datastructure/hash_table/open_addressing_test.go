package hash_table

import (
	"testing"

	"github.com/oinume/algo/src/datastructure/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAddressing_Put(t *testing.T) {
	a := assert.New(t)
	// TODO: table driven test
	hashTable := NewOpenAddresssing()
	a.Nil(hashTable.Put(&types.Object{1}, &types.Object{1}))
	a.Equal(1, hashTable.Size())
}

func TestOpenAddressing_Get(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	hashTable := NewOpenAddresssing()
	hashTable.Put(&types.Object{1}, &types.Object{10})
	hashTable.Put(&types.Object{2}, &types.Object{20})
	a.Equal(2, hashTable.Size())

	actual, err := hashTable.Get(&types.Object{2})
	r.NoError(err)
	a.Equal(&types.Object{20}, actual)
}
