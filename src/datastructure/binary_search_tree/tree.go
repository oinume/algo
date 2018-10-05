package binary_search_tree

import "fmt"

type Node struct {
	value int64
	left  *Node
	right *Node
}

func NewNode(value int64) *Node {
	return &Node{value: value}
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) Value() int64 {
	return n.value
}

func (n *Node) IsEqual(other *Node) bool {
	return n.Value() == other.Value()
}

func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

type Tree struct {
	root *Node
}

func New(root *Node) *Tree {
	return &Tree{root: root}
}

func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) Find(target int64) (*Node, error) {
	for n := t.Root(); n != nil; {
		v := n.Value()
		if v == target {
			return n, nil
		}
		if v > target {
			n = n.Left()
		} else {
			n = n.Right()
		}
	}
	return nil, ErrNotFound
}

/*
     5
   3   6

4を挿入したい

1. n.Value():5 > target:4 -> leftへ行く
2. n.Value():3 < target:4 -> rightがnil。insertToLeft=falseでforループを終了
3. 挿入処理: parent=3, insertToLeft=falseなのでnewNodeをparent.rightに挿入
*/
func (t *Tree) Insert(target int64) (*Node, error) {
	var parent *Node
	insertToLeft := false
	for n := t.Root(); n != nil; {
		if n.Value() == target {
			return nil, ErrAlreadyExists
		}
		if n.Value() > target {
			parent = n
			n = n.Left()
			insertToLeft = true
		} else { // n.Value() < target
			parent = n
			n = n.Right()
			insertToLeft = false
		}
	}

	newNode := NewNode(target)
	if parent == nil {
		t.root = newNode
		return newNode, nil
	}
	if insertToLeft {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	return newNode, nil
}

func (t *Tree) Remove(target int64) (*Node, error) {
	var parent *Node
	isLeftChild := false // The current node is left child of parent
	for n := t.root; n != nil; {
		if n.Value() == target { // Found
			if n.IsLeaf() { // The node is a leaf, just remove the node itself
				if parent == nil {
					t.root = nil
				} else if isLeftChild {
					parent.left = nil
				} else {
					parent.right = nil
				}
			} else if n.left == nil {
				// The node has only right child so replace this node with the right child
				if parent == nil {
					t.root = n.Right()
				} else if isLeftChild {
					parent.left = n.Right()
				} else {
					parent.right = n.Right()
				}
			} else if n.right == nil {
				if parent == nil {
					t.root = n.Left()
				} else if isLeftChild {
					parent.left = n.Left()
				} else {
					parent.right = n.Left()
				}
			} else {
				// Remove smallest value node in right subtree
				smallest := t.removeSmallestNode(n, n.Right())
				// Replace current node with the smallest node
				if parent == nil {
					t.root = smallest
				} else if isLeftChild {
					parent.left = smallest
				} else {
					parent.right = smallest
				}
				smallest.left = n.Left()
				smallest.right = n.Right()
			}
			// Successfully removed
			return n, nil
		} else if n.Value() > target {
			// Go to left child
			parent = n
			isLeftChild = true
			n = n.Left()
		} else {
			// Go to right child
			parent = n
			isLeftChild = false
			n = n.Right()
		}
	}
	return nil, ErrNotFound
}

func (t *Tree) removeSmallestNode(parent *Node, current *Node) *Node {
	isLeftChild := false
	for current.Left() != nil {
		parent = current
		isLeftChild = true
		current = current.Left()
	}

	if isLeftChild {
		parent.left = current.Right()
	} else {
		parent.right = current.Right()
	}
	return current
}

func (t *Tree) IsEqual(other *Tree) bool {
	// TODO: implement
	return true
}

var (
	ErrAlreadyExists = fmt.Errorf("already exists in this tree")
	ErrNotFound      = fmt.Errorf("not found in this tree")
)
