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
	name      string // For debugging
}

func NewNode(minDegree int, leaf bool) *Node {
	return &Node{
		minDegree: minDegree,
		leaf:      leaf,
		keys:      make([]int64, 0, minDegree*2),
		children:  make([]*Node, 0, minDegree*2),
	}
}

func (n *Node) IsLeaf() bool {
	return n.leaf
}

func (n *Node) Children() []*Node {
	return n.children
}

func (n *Node) insertNonFull(value int64) {
	if n.IsLeaf() {
		for i, key := range n.keys {
			if value < key {
				n.keys = insertKeyAt(n.keys, i, value)
				return
			}
		}
		n.keys = append(n.keys, value)
	} else {
		// TODO: add unit test
		i := n.locateSubtree(value)
		child := n.children[i]
		if child.needSplit() {
			n.splitChild(i, child)
			if value > n.keys[i] {
				child = n.children[i+1]
			}
		}
		child.insertNonFull(value)
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

func (n *Node) splitChild(index int, target *Node) {
	t := n.minDegree
	newChild := NewNode(n.minDegree, target.leaf)
	newChild.keys = target.keys[t:]
	newChild.name = "newChild"
	//for _, v := range target.keys[t:] {
	//	newChild.keys = append(newChild.keys, v)
	//}
	if !target.leaf {
		newChild.children = target.children[t:]
	}

	// Add newChild to this node(n)
	n.children = insertChildAt(n.children, index+1, newChild)
	n.keys = insertKeyAt(n.keys, index, target.keys[t-1])

	// Update target's keys and children for splitting
	target.keys = target.keys[0 : t-1]
	target.children = target.children[0:t]
	/*
	   def split_child(self, i, y):
	       t = self.t
	       z = BTree.Node(t)

	       z.is_leaf = y.is_leaf
	       z.keys = y.keys[t:]
	       if not y.is_leaf:
	           z.children = y.children[t:]

	       self.children.insert(i + 1, z)
	       self.keys.insert(i, y.keys[t - 1])

	       y.keys = y.keys[0:t - 1]
	       y.children = y.children[0:t]
	*/
}

func (n *Node) Dump(w io.Writer, pad int) {
	fmt.Fprintf(w, "%s%v:%s\n", strings.Repeat("-", pad), n.keys, n.name)
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

// TODO: bind to n
func insertKeyAt(keys []int64, position int, value int64) []int64 {
	s := make([]int64, 0, len(keys)+1)
	for i, v := range keys {
		if i == position {
			s = append(s, value)
		}
		s = append(s, v)
	}
	// TODO: confirm
	if position >= len(keys) {
		s = append(s, value)
	}
	return s
}

func insertChildAt(children []*Node, position int, value *Node) []*Node {
	s := make([]*Node, 0, len(children)+1)
	for i, v := range children {
		if i == position {
			s = append(s, value)
		}
		s = append(s, v)
	}
	if position >= len(children) {
		s = append(s, value)
	}
	return s
}
