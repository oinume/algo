package binary_search_tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/oinume/algo/testings"
)

var cmpOptions = []cmp.Option{
	cmp.AllowUnexported(Node{}, Tree{}),
}

func TestNewTree(t *testing.T) {
	root := NewNode(100)
	tree := NewTree(NewNode(100))
	testings.AssertEqual(t, root, tree.Root(), "unexpected tree root", cmpOptions...)
}

func TestTree_Find(t *testing.T) {
	tree := NewTree(NewNode(5))
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
			if err != nil {
				t.Fatalf("tree.Find returns unexpected error: %v", err)
			}
			testings.AssertEqual(t, test.want, got, "tree.Find returns unexpected node", cmpOptions...)
		}
	})

	t.Run("not found", func(t *testing.T) {
		_, err := tree.Find(100)
		testings.AssertEqual(t, ErrNotFound.Error(), err.Error(), "err must be ErrNotFound")
	})
}

func TestTree_Insert(t *testing.T) {
	tests := []struct {
		insert       int64
		want         *Node
		wantTreeFunc func() *Tree
	}{
		{
			insert: 5,
			want:   NewNode(5),
			wantTreeFunc: func() *Tree {
				return NewTree(NewNode(5))
			},
		},
		{
			insert: 3,
			want:   NewNode(3),
			wantTreeFunc: func() *Tree {
				tree := NewTree(NewNode(5))
				tree.root.left = NewNode(3)
				return tree
			},
		},
		{
			insert: 6,
			want:   NewNode(6),
			wantTreeFunc: func() *Tree {
				tree := NewTree(NewNode(5))
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
				tree := NewTree(NewNode(5))
				left := NewNode(3)
				tree.root.left = left
				left.right = NewNode(4)
				tree.root.right = NewNode(6)
				return tree
			},
		},
	}

	tree := NewTree(nil)
	for _, test := range tests {
		node, err := tree.Insert(test.insert)
		if err != nil {
			t.Fatalf("tree.Insert failed: %v", err)
		}
		testings.AssertEqual(t, test.want.Value(), node.Value(), "")
		testings.AssertEqual(t, test.wantTreeFunc(), tree, "", cmpOptions...)
	}
}

func TestTree_Insert_Exist(t *testing.T) {
	tree := NewTree(NewNode(5))
	tree.root.left = NewNode(3)
	tree.root.right = NewNode(6)

	_, err := tree.Insert(6)
	testings.AssertEqual(t, ErrAlreadyExists.Error(), err.Error(), "err must be ErrAlreadyExists")
}

func TestTree_Remove(t *testing.T) {
	tests := map[string]struct {
		setupTree        func() *Tree
		target           int64
		removedAfterTree func() *Tree
	}{
		"remove leaf": {
			//     5
			//  ／    ＼
			// 3       6
			//  ＼
			//    4
			setupTree: func() *Tree {
				tree := NewTree(NewNode(5))
				left := NewNode(3)
				tree.root.left = left
				left.right = NewNode(4)
				tree.root.right = NewNode(6)
				return tree
			},
			target: 4,
			removedAfterTree: func() *Tree {
				tree := NewTree(NewNode(5))
				tree.root.left = NewNode(3)
				tree.root.right = NewNode(6)
				return tree
			},
		},
		"remove node with left child": {
			//          9
			//        ／  ＼
			//       5     14
			//     ／
			//    3
			//  ／  ＼
			// 1     4
			setupTree: func() *Tree {
				tree := NewTree(NewNode(9))
				tree.root.left = NewNode(5)
				tree.root.right = NewNode(14)
				left2 := NewNode(3)
				left2.left = NewNode(1)
				left2.right = NewNode(4)
				tree.root.left.left = left2
				return tree
			},
			target: 5,
			removedAfterTree: func() *Tree {
				tree := NewTree(NewNode(9))
				left := NewNode(3)
				left.left = NewNode(1)
				left.right = NewNode(4)
				tree.root.left = left
				tree.root.right = NewNode(14)
				return tree
			},
		},
		"remove node with right child": {
			//       9
			//     ／  ＼
			//    2     10
			//            ＼
			//             17
			//           ／  ＼
			//          13    19
			setupTree: func() *Tree {
				tree := NewTree(NewNode(9))
				tree.root.left = NewNode(2)
				tree.root.right = NewNode(10)
				right2 := NewNode(17)
				right2.left = NewNode(13)
				right2.right = NewNode(19)
				tree.root.right.right = right2
				return tree
			},
			target: 10,
			removedAfterTree: func() *Tree {
				tree := NewTree(NewNode(9))
				tree.root.left = NewNode(2)
				right := NewNode(17)
				right.left = NewNode(13)
				right.right = NewNode(19)
				tree.root.right = right
				return tree
			},
		},
		"remove node with left and right child": {
			//           20
			//         ／   ＼
			//        7      23
			//     ／   ＼     ＼
			//    4      18     29
			//  ／  ＼   ／
			// 2     5  10
			//            ＼
			//             15
			setupTree: func() *Tree {
				tree := NewTree(NewNode(20))
				tree.root.left = NewNode(7)
				tree.root.right = NewNode(23)

				left7 := NewNode(4)
				left7.left = NewNode(2)
				left7.right = NewNode(5)
				tree.root.left.left = left7

				right7 := NewNode(18)
				right7.left = NewNode(10)
				right7.left.right = NewNode(15)
				tree.root.left.right = right7

				tree.root.right.right = NewNode(29)

				return tree
			},
			target: 7,
			//           20
			//         ／   ＼
			//        10     23
			//     ／   ＼     ＼
			//    4      18     29
			//  ／  ＼   ／
			// 2     5  15
			removedAfterTree: func() *Tree {
				tree := NewTree(NewNode(20))
				tree.root.left = NewNode(10)
				tree.root.right = NewNode(23)

				left10 := NewNode(4)
				left10.left = NewNode(2)
				left10.right = NewNode(5)
				tree.root.left.left = left10

				right10 := NewNode(18)
				right10.left = NewNode(15)
				tree.root.left.right = right10

				tree.root.right.right = NewNode(29)

				return tree
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tree := test.setupTree()
			got, err := tree.Remove(test.target)
			if err != nil {
				t.Fatalf("target is %v", test.target)
			}
			testings.AssertEqual(t, test.target, got.Value(), "", cmpOptions...)
			testings.AssertEqual(t, test.removedAfterTree(), tree, "", cmpOptions...)
		})
	}
}

func TestTree_Remove_NotFound(t *testing.T) {
	tree := NewTree(NewNode(5))
	tree.root.left = NewNode(3)
	tree.root.right = NewNode(6)
	tree.root.right.right = NewNode(10)

	_, err := tree.Remove(15)
	testings.AssertEqual(t, ErrNotFound.Error(), err.Error(), "err must be ErrNotFound")
}
