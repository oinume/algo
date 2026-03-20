package graph

type Visitor[T comparable] interface {
	Visit(g *Graph[T], v *Node[T])
}

type nopVisitor[T comparable] struct{}

func (nv *nopVisitor[T]) Visit(g *Graph[T], v *Node[T]) {}

func NewListVisitor[T comparable]() *listVisitor[T] {
	return &listVisitor[T]{
		list: make([]*Node[T], 0, 100),
	}
}

type listVisitor[T comparable] struct {
	list []*Node[T]
}

func (lv *listVisitor[T]) Visit(g *Graph[T], v *Node[T]) {
	lv.list = append(lv.list, v)
}
