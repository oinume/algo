package graph

import "fmt"

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

func (v *Vertex) String() string {
	return v.value
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

func (e *Edge) String() string {
	return fmt.Sprintf("[%v -> %v]", e.start, e.end)
}

type Graph struct {
	// All vertices this Graph has
	vertices *vertexSet
	// Edges per vertex
	edges map[*Vertex][]*Edge
}

func New() *Graph {
	return &Graph{
		vertices: NewVertexSet(100),
		edges:    make(map[*Vertex][]*Edge, 100),
	}
}

func (g *Graph) Edges(v *Vertex) []*Edge {
	if edges, ok := g.edges[v]; ok {
		return edges
	}
	return nil
}

func (g *Graph) Vertices() []*Vertex {
	return g.vertices.Values()
}

func (g *Graph) AddVertexWithEdges(v *Vertex, edges []*Edge) {
	g.edges[v] = edges
	g.vertices.Add(v)
}

func (g *Graph) AddEdge(start *Vertex, end *Vertex) {
	g.edges[start] = append(g.edges[start], &Edge{start: start, end: end})
	g.vertices.Add(start)
}
