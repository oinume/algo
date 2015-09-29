package data

type VisitFunc func(value *Object)

type BinaryTreeNode struct {
	Value     *Object
	Left      *BinaryTreeNode
	Right     *BinaryTreeNode
	visitFunc VisitFunc
}

func (n *BinaryTreeNode) Accept(visitor VisitFunc) {
	visitor(n.Value)
	if n.Left != nil {
		n.Left.Accept(visitor)
	}
	if n.Right != nil {
		n.Right.Accept(visitor)
	}
}
