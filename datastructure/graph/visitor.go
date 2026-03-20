package graph

type Visitor interface {
	Visit(g *Graph, v *Node)
}

type nopVisitor struct{}

func (nv *nopVisitor) Visit(g *Graph, v *Node) {}

func (nv *nopVisitor) Visited() []*Node {
	return nil
}

func NewListVisitor() Visitor {
	return &listVisitor{
		list: make([]*Node, 0, 100),
	}
}

type listVisitor struct {
	list []*Node
}

func (lv *listVisitor) Visit(g *Graph, v *Node) {
	lv.list = append(lv.list, v)
}
