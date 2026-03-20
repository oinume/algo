package graph

import (
	"bytes"
	"fmt"
)

type Node[T comparable] struct {
	value T
}

func NewNode[T comparable](v T) *Node[T] {
	return &Node[T]{value: v}
}

func (v *Node[T]) IsEqual(other *Node[T]) bool {
	return v.value == other.value || v == other
}

func (v *Node[T]) String() string {
	return fmt.Sprint(v.value)
}

type Edge[T comparable] struct {
	start *Node[T]
	end   *Node[T]
}

func newEdge[T comparable](start, end *Node[T]) *Edge[T] {
	return &Edge[T]{
		start: start,
		end:   end,
	}
}

func (e *Edge[T]) String() string {
	return fmt.Sprintf("[%v -> %v]", e.start, e.end)
}

type Graph[T comparable] struct {
	// All nodes this Graph has
	nodes *nodeSet[T]
	// Edges per node
	edges map[*Node[T]][]*Edge[T]
}

func New[T comparable]() *Graph[T] {
	return &Graph[T]{
		nodes: newNodeSet[T](100),
		edges: make(map[*Node[T]][]*Edge[T], 100),
	}
}

func (g *Graph[T]) Edges(v *Node[T]) []*Edge[T] {
	if edges, ok := g.edges[v]; ok {
		return edges
	}
	return nil
}

func (g *Graph[T]) Dump() string {
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

func (g *Graph[T]) Nodes() []*Node[T] {
	return g.nodes.Values()
}

func (g *Graph[T]) AddEdge(start *Node[T], end *Node[T]) {
	g.edges[start] = append(g.edges[start], newEdge(start, end))
	g.nodes.Add(start)
}
