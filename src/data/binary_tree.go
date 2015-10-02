package data

type Visitor interface {
	Visit(node *BinaryTreeNode, depth int)
}

type BinaryTreeNode struct {
	Value *Object
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (n *BinaryTreeNode) Accept(visitor Visitor, depth int) {
	visitor.Visit(n, depth)
	if n.Left != nil {
		n.Left.Accept(visitor, depth+1)
	}
	if n.Right != nil {
		n.Right.Accept(visitor, depth+1)
	}
}
