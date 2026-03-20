package graph

import (
	"bytes"
	"fmt"
)

type Node struct {
	value string
}

func NewNode(v string) *Node {
	return &Node{value: v}
}

func (v *Node) IsEqual(other *Node) bool {
	if v.value != "" && other.value != "" && v.value == other.value {
		return true
	}
	if v == other {
		return true
	}
	return false
}

func (v *Node) String() string {
	return v.value
}

type Edge struct {
	start *Node
	end   *Node
}

func newEdge(start, end *Node) *Edge {
	return &Edge{
		start: start,
		end:   end,
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("[%v -> %v]", e.start, e.end)
}

type Graph struct {
	// All nodes this Graph has
	nodes *nodeSet
	// Edges per node
	edges map[*Node][]*Edge
}

func New() *Graph {
	return &Graph{
		nodes: NewNodeSet(100),
		edges: make(map[*Node][]*Edge, 100),
	}
}

func (g *Graph) Edges(v *Node) []*Edge {
	if edges, ok := g.edges[v]; ok {
		return edges
	}
	return nil
}

func (g *Graph) Dump() string {
	b := new(bytes.Buffer)
	for _, v := range g.Nodes() {
		edges := g.Edges(v)
		if edges == nil {
			continue
		}
		for _, e := range edges {
			_, _ = fmt.Fprintf(b, "%v:", v.String())
			_, _ = fmt.Fprintln(b, e.String())
		}
	}
	return b.String()
}

func (g *Graph) Nodes() []*Node {
	return g.nodes.Values()
}

func (g *Graph) AddEdge(start *Node, end *Node) {
	g.edges[start] = append(g.edges[start], newEdge(start, end))
	g.nodes.Add(start)
}
