package graph

import (
	"fmt"
	"testing"
)

// https://www.geeksforgeeks.org/graph-data-structure-and-algorithms/
func Test_dfsFinder_Find(t *testing.T) {
	graph := New()
	v1 := NewVertex("1")
	v2 := NewVertex("2")
	e1_2 := NewEdge(v1, v2)
	graph.AddVertexWithEdges(v1, []*Edge{e1_2})

	f := NewDFSFinder()
	if found := f.Find(graph, v1, v2, &nopVisitor{}); !found {
		t.Errorf("v1 not found somehow")
	}
}

/*
　　　A
　　／ ＼
　B　ー　C
／
D
*/
func newGraph() *Graph {
	g := New()
	v := make([]*Vertex, 4)
	for i := 0; i < len(v); i++ {
		v[i] = NewVertex(fmt.Sprint(i))
	}
	edges := map[*Vertex][]*Edge{
		v[0] /* A */ : {NewEdge(v[0], v[1]), NewEdge(v[0], v[2])},
		v[1] /* B */ : {NewEdge(v[1], v[2]), NewEdge(v[1], v[3])},
		v[2] /* C */ : {NewEdge(v[2], v[0]), NewEdge(v[2], v[1])},
		v[3] /* D */ : {NewEdge(v[3], v[1])},
	}
	for v, e := range edges {
		g.AddVertexWithEdges(v, e)
	}
	return g
}
