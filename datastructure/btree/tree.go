// http://d.hatena.ne.jp/naoya/20090412/btree
package btree

import "fmt"

var ErrAlreadyExists = fmt.Errorf("the key already exists")

type Tree struct {
	root Node
}

func NewTree(order int) *Tree {
	return &Tree{
		root: NewNode(order),
	}
}

/*
class BTree:
    def __init__(self, t=2):
        self.t = t
        self.root = BTree.Node(t)
        self.root.is_leaf = True

    def insert(self, k):
        r = self.root
        if len(r) == 2 * self.t - 1:
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
