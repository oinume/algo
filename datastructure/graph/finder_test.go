package graph

import (
	"fmt"
	"reflect"
	"testing"
)

// https://www.geeksforgeeks.org/graph-data-structure-and-algorithms/
func Test_dfsFinder_Find_2V(t *testing.T) {
	graph := New[string]()
	v1 := NewNode("1")
	v2 := NewNode("2")
	graph.AddEdge(v1, v2)

	tests := map[string]struct {
		finder Finder[string]
	}{
		"dfsRecursiveFinder": {
			finder: NewDFSRecursiveFinder[string](),
		},
		"dfsLoopFinder": {
			finder: NewDFSLoopFinder[string](),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if found := test.finder.Find(graph, v1, v2, &nopVisitor[string]{}); !found {
				t.Errorf("v1 not found somehow")
			}
		})
	}
}

// https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/
func Test_dfsFinder_Find_4V(t *testing.T) {
	v := make([]*Node[string], 4)
	for i := 0; i < len(v); i++ {
		v[i] = NewNode(fmt.Sprint(i))
	}
	graph := New[string]()
	graph.AddEdge(v[0], v[1])
	graph.AddEdge(v[0], v[2])
	graph.AddEdge(v[1], v[2])
	graph.AddEdge(v[2], v[0])
	graph.AddEdge(v[2], v[3])
	graph.AddEdge(v[3], v[3])

	tests := map[string]struct {
		finder           Finder[string]
		wantVisitedNodes []*Node[string]
	}{
		"dfsRecursiveFinder": {
			finder: NewDFSRecursiveFinder[string](),
			wantVisitedNodes: []*Node[string]{
				NewNode("2"),
				NewNode("0"),
				NewNode("1"),
				NewNode("3"),
			},
		},
		"dfsLoopFinder": {
			finder: NewDFSLoopFinder[string](),
			wantVisitedNodes: []*Node[string]{
				NewNode("2"),
				NewNode("3"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lv := &listVisitor[string]{}
			if found := test.finder.Find(graph, v[2], v[3], lv); !found {
				t.Errorf("%v not found somehow", v[3])
			}
			if got, want := lv.list, test.wantVisitedNodes; !reflect.DeepEqual(got, want) {
				t.Errorf("visiting node is not correct: got = %+v, want = %+v", got, want)
			}
		})
	}
}
