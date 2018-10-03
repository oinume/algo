package binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	root := NewNode(100)
	tree := New(NewNode(100))
	a.Equal(tree.Root(), root)
}

func TestTree_Insert(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	tests := []struct {
		insert       int64
		want         *Node
		wantTreeFunc func() *Tree
	}{
		{
			insert: 5,
			want:   NewNode(5),
			wantTreeFunc: func() *Tree {
				return New(NewNode(5))
			},
		},
		{
			insert: 3,
			want:   NewNode(3),
			wantTreeFunc: func() *Tree {
				tree := New(NewNode(5))
				tree.root.left = NewNode(3)
				return tree
			},
		},
		{
			insert: 6,
			want:   NewNode(6),
			wantTreeFunc: func() *Tree {
				tree := New(NewNode(5))
				tree.root.left = NewNode(3)
				tree.root.right = NewNode(6)
				return tree
			},
		},
		//     5
		//  ／    ＼
		// 3       6
		//  ＼
		//    4
		{
			insert: 4,
			want:   NewNode(4),
			wantTreeFunc: func() *Tree {
				tree := New(NewNode(5))
				left := NewNode(3)
				tree.root.left = left
				left.right = NewNode(4)
				tree.root.right = NewNode(6)
				return tree
			},
		},
	}

	tree := New(nil)
	for _, test := range tests {
		node, err := tree.Insert(test.insert)
		if err != nil {
			r.NoError(err, "tree.Insert failed")
		}
		a.Equal(test.want.Value(), node.Value())
		a.Equal(test.wantTreeFunc(), tree)
	}
}
