package avl_tree

type Tree struct {
	root *Node
}

func NewTree(root *Node) *Tree {
	return &Tree{root: root}
}

func (t *Tree) Insert(value int64) (*Node, error) {
	return nil, nil
}
