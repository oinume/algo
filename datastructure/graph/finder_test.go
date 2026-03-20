package graph

import (
	"fmt"
	"reflect"
	"testing"
)

// https://www.geeksforgeeks.org/graph-data-structure-and-algorithms/
func Test_dfsFinder_Find_2N(t *testing.T) {
	graph := New[string]()
	n1 := NewNode("1")
	n2 := NewNode("2")
	graph.AddEdge(n1, n2)

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
			if found := test.finder.Find(graph, n1, n2, &nopVisitor[string]{}); !found {
				t.Errorf("n1 not found somehow")
			}
		})
	}
}

// https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/
func Test_dfsFinder_Find_4N(t *testing.T) {
	nodes := make([]*Node[string], 4)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = NewNode(fmt.Sprint(i))
	}
	graph := New[string]()
	graph.AddEdge(nodes[0], nodes[1])
	graph.AddEdge(nodes[0], nodes[2])
	graph.AddEdge(nodes[1], nodes[2])
	graph.AddEdge(nodes[2], nodes[0])
	graph.AddEdge(nodes[2], nodes[3])
	graph.AddEdge(nodes[3], nodes[3])

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
			if found := test.finder.Find(graph, nodes[2], nodes[3], lv); !found {
				t.Errorf("%v not found somehow", nodes[3])
			}
			if got, want := lv.list, test.wantVisitedNodes; !reflect.DeepEqual(got, want) {
				t.Errorf("visiting node is not correct: got = %+v, want = %+v", got, want)
			}
		})
	}
}
