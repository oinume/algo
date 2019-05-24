package graph

type Visitor interface {
	Visit(g *Graph, v *Vertex)
}

type nopVisitor struct{}

func (nv *nopVisitor) Visit(g *Graph, v *Vertex) {}

func (nv *nopVisitor) Visited() []*Vertex {
	return nil
}

func NewListVisitor() Visitor {
	return &listVisitor{
		list: make([]*Vertex, 0, 100),
	}
}

type listVisitor struct {
	list []*Vertex
}

func (lv *listVisitor) Visit(g *Graph, v *Vertex) {
	lv.list = append(lv.list, v)
}
