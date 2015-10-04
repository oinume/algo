package data

// Binary tree node
// Left: left child
// Right: right child
type BinaryTreeNode struct {
	Value *Object
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// Visitor interface for binary tree
type BinaryTreeVisitor interface {
	Visit(node *BinaryTreeNode, depth int)
}

// Traverse from this node by pre-order
func (n *BinaryTreeNode) TraversePreorder(v BinaryTreeVisitor, depth int) {
	v.Visit(n, depth)
	if n.Left != nil {
		n.Left.TraversePreorder(v, depth+1)
	}
	if n.Right != nil {
		n.Right.TraversePreorder(v, depth+1)
	}
}

// Traverse from this node by in-order
func (n *BinaryTreeNode) TraverseInorder(v BinaryTreeVisitor, depth int) {
	if n.Left != nil {
		n.Left.TraverseInorder(v, depth+1)
	}
	v.Visit(n, depth)
	if n.Right != nil {
		n.Right.TraverseInorder(v, depth+1)
	}
}

// Traverse from this node by post-order
func (n *BinaryTreeNode) TraversePostorder(v BinaryTreeVisitor, depth int) {
	if n.Left != nil {
		n.Left.TraversePostorder(v, depth+1)
	}
	if n.Right != nil {
		n.Right.TraversePostorder(v, depth+1)
	}
	v.Visit(n, depth)
}
