package data

type Visitor interface {
	Visit(node *BinaryTreeNode, depth int)
}

type BinaryTreeNode struct {
	Value *Object
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (n *BinaryTreeNode) TraversePreorder(v Visitor, depth int) {
	v.Visit(n, depth)
	if n.Left != nil {
		n.Left.TraversePreorder(v, depth+1)
	}
	if n.Right != nil {
		n.Right.TraversePreorder(v, depth+1)
	}
}

func (n *BinaryTreeNode) TraverseInorder(v Visitor, depth int) {
	if n.Left != nil {
		n.Left.TraverseInorder(v, depth+1)
	}
	v.Visit(n, depth)
	if n.Right != nil {
		n.Right.TraverseInorder(v, depth+1)
	}
}

func (n *BinaryTreeNode) TraversePostorder(v Visitor, depth int) {
	if n.Left != nil {
		n.Left.TraversePostorder(v, depth+1)
	}
	if n.Right != nil {
		n.Right.TraversePostorder(v, depth+1)
	}
	v.Visit(n, depth)
}
