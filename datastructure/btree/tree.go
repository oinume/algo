package btree

import (
	"fmt"
	"io"
)

var ErrAlreadyExists = fmt.Errorf("the key already exists")

type Tree struct {
	root      *Node
	minDegree int
}

func NewTree(minDegree int) *Tree {
	return &Tree{
		root:      NewNode(minDegree, true),
		minDegree: minDegree,
	}
}

func (t *Tree) NewNode() *Node {
	return NewNode(t.minDegree, false)
}

func (t *Tree) Insert(value int64) {
	r := t.root
	if r.needSplit() {
		n := NewNode(t.minDegree, false)
		n.name = "n"
		n.children = append(n.children, r)
		r.name = "r"
		n.splitChild(0, r)
		n.insertNonFull(value)
		t.root = n
	} else {
		r.insertNonFull(value)
	}
}

func (t *Tree) Dump(w io.Writer) {
	t.root.Dump(w, 1)
}
