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

func TestTree_Find(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	tree := New(NewNode(5))
	left := NewNode(3)
	tree.root.left = left
	left.right = NewNode(4)
	tree.root.right = NewNode(6)

	t.Run("normal", func(t *testing.T) {
		tests := []struct {
			input int64
			want  *Node
		}{
			{input: 4, want: NewNode(4)},
		}
		for _, test := range tests {
			got, err := tree.Find(test.input)
			r.NoError(err)
			a.Equal(test.want, got)
		}
	})

	t.Run("not found", func(t *testing.T) {
		_, err := tree.Find(100)
		r.Equal(ErrNotFound, err)
	})
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

func TestTree_Insert_Exist(t *testing.T) {
	r := require.New(t)
	tree := New(NewNode(5))
	tree.root.left = NewNode(3)
	tree.root.right = NewNode(6)

	_, err := tree.Insert(6)
	r.Equal(ErrAlreadyExists, err)
}

func TestTree_Remove(t *testing.T) {
	//a := assert.New(t)
	//r := require.New(t)
}
