package binary_search_tree

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

type Tree struct {
	root *Node
}

func New(root *Node) *Tree {
	return &Tree{root: root}
}

func (t *Tree) Root() *Node {
	return t.root
}
