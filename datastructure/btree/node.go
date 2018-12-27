package btree

import "io"

type Node struct {
	minDegree int
	leaf      bool
	keys      []int64
}

func NewNode(minDegree int, leaf bool) *Node {
	return &Node{
		minDegree: minDegree,
		leaf:      leaf,
		keys:      nil,
	}
}

func (n *Node) IsLeaf() bool {
	return n.leaf
}

func (n *Node) Children() []*Node {
	return nil
}

func (n *Node) InsertNonFull(value int64) {
	if n.IsLeaf() {
		for i, key := range n.keys {
			if value < key {
				// Insert value at i
				n.keys = append(n.keys[:i], append([]int64{value}, n.keys[i:]...)...)
				return
			}
		}
		n.keys = append(n.keys, value)
	} else {
		// TODO
	}
	/*
		def insert_nonfull(self, k):
			if self.is_leaf:
				i = 0
				for i in xrange(len(self)):
					if k < self.keys[i]:
						self.keys.insert(i, k)
						return self
				self.keys.append(k)
			else:
				i = self.locate_subtree(k)
				c = self.children[i]
				if (len(c) == 2 * self.t - 1):
					self.split_child(i, c)
					if k > self.keys[i]:
						c = self.children[i + 1]
				c.insert_nonfull(k)
	*/
}

func (n *Node) Dump(w io.Writer) {

}
