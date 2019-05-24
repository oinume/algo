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
	graph.AddEdge(v1, v2)

	tests := map[string]struct {
		finder Finder
	}{
		"dfsRecursiveFinder": {
			finder: NewDFSRecursiveFinder(),
		},
		"dfsLoopFinder": {
			finder: NewDFSLoopFinder(),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if found := test.finder.Find(graph, v1, v2, &nopVisitor{}); !found {
				t.Errorf("v1 not found somehow")
			}
		})
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

	tests := map[string]struct {
		finder              Finder
		wantVisitedVertices []*Vertex
	}{
		"dfsRecursiveFinder": {
			finder: NewDFSRecursiveFinder(),
			wantVisitedVertices: []*Vertex{
				NewVertex("2"),
				NewVertex("0"),
				NewVertex("1"),
				NewVertex("3"),
			},
		},
		"dfsLoopFinder": {
			finder: NewDFSLoopFinder(),
			wantVisitedVertices: []*Vertex{
				NewVertex("2"),
				NewVertex("3"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lv := &listVisitor{}
			if found := test.finder.Find(graph, v[2], v[3], lv); !found {
				t.Errorf("%v not found somehow", v[3])
			}
			if got, want := lv.list, test.wantVisitedVertices; !reflect.DeepEqual(got, want) {
				t.Errorf("visiting vertex is not correct: got = %+v, want = %+v", got, want)
			}
		})
	}
}
