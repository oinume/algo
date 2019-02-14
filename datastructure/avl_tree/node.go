package avl_tree

type Node struct {
	value int64
	left  *Node
	right *Node
	// Positive value if right child's height is greater than left
	// Negative value if left child's height is greater than right
	// Zero if left and right height are same or the node is leaf
	balance int
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
