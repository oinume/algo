package graph

import (
	"fmt"
	"reflect"
	"testing"
)

// https://www.geeksforgeeks.org/graph-data-structure-and-algorithms/
func Test_dfsFinder_Find_2V(t *testing.T) {
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

// https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/
func Test_dfsFinder_Find_4V(t *testing.T) {
	v := make([]*Vertex, 4)
	for i := 0; i < len(v); i++ {
		v[i] = NewVertex(fmt.Sprint(i))
	}
	graph := New()
	graph.AddEdge(v[0], v[1])
	graph.AddEdge(v[0], v[2])
	graph.AddEdge(v[1], v[2])
	graph.AddEdge(v[2], v[0])
	graph.AddEdge(v[2], v[3])
	graph.AddEdge(v[3], v[3])

	f := NewDFSFinder()
	lv := &listVisitor{}
	if found := f.Find(graph, v[2], v[3], lv); !found {
		t.Errorf("v3 not found somehow")
	}

	want := []*Vertex{
		NewVertex("2"),
		NewVertex("0"),
		NewVertex("1"),
		NewVertex("3"),
	}
	if !reflect.DeepEqual(lv.list, want) {
		t.Errorf("visiting vertex is not correct: got = %+v, want = %+v", lv.list, want)
	}
}
