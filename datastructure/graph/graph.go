package graph

type Vertex struct {
	value string
}

func NewVertex(v string) *Vertex {
	return &Vertex{value: v}
}

func (v *Vertex) IsEqual(other *Vertex) bool {
	if v.value != "" && other.value != "" && v.value == other.value {
		return true
	}
	if v == other {
		return true
	}
	return false
}

type Edge struct {
	start *Vertex
	end   *Vertex
}

func NewEdge(start, end *Vertex) *Edge {
	return &Edge{
		start: start,
		end:   end,
	}
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

func (g *Graph) Vertexes() []*Vertex {
	return g.vertexes.Values()
}

func (g *Graph) AddVertexWithEdges(v *Vertex, edges []*Edge) {
	g.edges[v] = edges
	g.vertexes.Add(v)
}
