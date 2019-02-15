package btree

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	minDegree int
	leaf      bool
	keys      []int64
	children  []*Node
}

func NewNode(minDegree int, leaf bool) *Node {
	return &Node{
		minDegree: minDegree,
		leaf:      leaf,
		keys:      nil,
		children:  nil, // TODO: make
	}
}

func (n *Node) IsLeaf() bool {
	return n.leaf
}

func (n *Node) Children() []*Node {
	return n.children
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
		i := n.locateSubtree(value)
		child := n.children[i]
		if n.needSplit() {
			n.splitChild()
			if value > n.keys[i] {
				child = n.children[i+1]
			}
		}
		child.InsertNonFull(value)
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

func (n *Node) needSplit() bool {
	return len(n.keys) == 2*n.minDegree-1
}

func (n *Node) locateSubtree(value int64) int {
	i := 0
	for ; i < len(n.keys); i++ {
		if value < n.keys[i] {
			return i
		}
	}
	return i
	/*
	   def locate_subtree(self, k):
	       i = 0
	       while (i < len(self)):
	           if k < self.keys[i]:
	               return i
	           i += 1
	       return i

	*/
}

func (n *Node) splitChild() {}

func (n *Node) Dump(w io.Writer, pad int) {
	fmt.Fprintf(w, "%s:%v\n", strings.Repeat("-", pad), n.keys)
	if n.leaf {
		return
	}
	for _, c := range n.children {
		c.Dump(w, pad+1)
	}
	/*
			        def show(self, pad):
		            print "%s%s" % ('-' * pad, self.keys)
		            if self.is_leaf:
		                return
		            else:
		                for c in self.children:
		                    c.show(pad + 1)

	*/
}
