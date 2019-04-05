package graph

import "testing"

// https://www.geeksforgeeks.org/graph-data-structure-and-algorithms/
func Test_dfsFinder_Find(t *testing.T) {
	graph := New()
	v1 := NewVertex("1")
	v2 := NewVertex("2")
	e1_2 := NewEdge(v1, v2)
	graph.AddVertexWithEdges(v1, []*Edge{e1_2})

	f := NewDFSFinder()
	if found := f.Find(graph, v1, v2); !found {
		t.Errorf("v1 not found somehow")
	}
}
