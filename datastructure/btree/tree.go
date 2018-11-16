package btree

type Tree struct {
	root Node
}

func NewTree(root Node) *Tree {
	return &Tree{
		root: root,
	}
}
