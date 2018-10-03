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

type Tree struct {
	root *Node
}

func New(root *Node) *Tree {
	return &Tree{root: root}
}

func (t *Tree) Root() *Node {
	return t.root
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
	//n := t.Root()
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

func (t *Tree) IsEqual(other *Tree) bool {
	// TODO: implement
	return true
}

var ErrAlreadyExists = fmt.Errorf("already exists in this tree")
