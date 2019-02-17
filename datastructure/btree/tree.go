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
		// TODO: unit test
		n := NewNode(t.minDegree, false)
		n.name = "n"
		n.children = append(n.children, r)
		//fmt.Printf("n.children = %+v\n", len(n.children))
		//println("---Insert---")
		//n.Dump(os.Stdout, 0)
		r.name = "r"
		n.splitChild(0, r)
		//t.Dump(os.Stdout)
		n.insertNonFull(value)
		t.root = n
		/*
		   s = BTree.Node(self.t)
		   s.children.append(r)
		   s.split_child(0, r)
		   s.insert_nonfull(k)
		   self.root = s
		*/
	} else {
		r.insertNonFull(value)
	}
}

func (t *Tree) Dump(w io.Writer) {
	t.root.Dump(w, 1)
}

/*
class BTree:
    def __init__(self, t=2):
        self.t = t # t is minimum degree
        self.root = BTree.Node(t)
        self.root.is_leaf = True

    def insert(self, k):
        r = self.root
        if len(r) == 2 * self.t - 1: #
            s = BTree.Node(self.t)
            s.children.append(r)
            s.split_child(0, r)
            s.insert_nonfull(k)
            self.root = s
        else:
            r.insert_nonfull(k)

    def delete(self, k):
        r = self.root
        if r.search(k) is None:
            return
        r.delete(k)
        if len(r) == 0:
            self.root = r.children[0]

    def search(self, k):
        return self.root.search(k)

    def show(self):
        self.root.show(1)
*/
