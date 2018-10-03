package binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	root := NewNode(100)
	tree := New(NewNode(100))
	a := assert.New(t)
	a.Equal(tree.Root(), root)
}
