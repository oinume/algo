package graph

type Vertex struct{}

type Edge struct {
	start Vertex
	end   Vertex
}

type Graph struct{}

func New() *Graph {
	return &Graph{}
}

func (g *Graph) Edges(v *Vertex) []*Edge {
	return nil
}
