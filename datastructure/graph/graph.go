package graph

type Vertex struct{}

func (v *Vertex) IsEqual(other *Vertex) bool {
	return v == other
}

type Edge struct {
	start *Vertex
	end   *Vertex
}

type Graph struct {
	// All vertexes this Graph has
	vertexes *vertexSet
	// Edges per vertex
	edges map[*Vertex][]*Edge
}

func New() *Graph {
	return &Graph{
		vertexes: NewVertexSet(100),
		edges:    make(map[*Vertex][]*Edge, 100),
	}
}

func (g *Graph) Edges(v *Vertex) []*Edge {
	if edges, ok := g.edges[v]; ok {
		return edges
	}
	return nil
}
