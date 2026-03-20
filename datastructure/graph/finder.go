package graph

import (
	"github.com/oinume/algo/datastructure/stack"
)

type Finder interface {
	Find(g *Graph, start *Node, target *Node, visitor Visitor) bool
}

// dfsRecursiveFinder is depth first search finder
type dfsRecursiveFinder struct {
	visited map[*Node]struct{}
}

func NewDFSRecursiveFinder() Finder {
	return &dfsRecursiveFinder{
		visited: make(map[*Node]struct{}, 100),
	}
}

func (dfs *dfsRecursiveFinder) Find(g *Graph, start *Node, target *Node, visitor Visitor) bool {
	//fmt.Printf("Find(): start = %+v\n", start)
	visitor.Visit(g, start)

	if start.IsEqual(target) {
		return true
	}
	if _, visited := dfs.visited[start]; visited {
		return false
	}

	dfs.visited[start] = struct{}{}
	edges := g.Edges(start)
	//fmt.Printf("edges = %+v\n", edges)
	for _, edge := range edges {
		if _, visited := dfs.visited[edge.end]; visited {
			continue
		}
		if result := dfs.Find(g, edge.end, target, visitor); result {
			return result
		}
	}

	return false
}

type dfsLoopFinder struct {
	visited map[*Node]struct{}
}

func NewDFSLoopFinder() Finder {
	return &dfsLoopFinder{
		visited: make(map[*Node]struct{}, 100),
	}
}

func (dfs *dfsLoopFinder) Find(g *Graph, start *Node, target *Node, visitor Visitor) bool {
	st := stack.New[*Node](g.nodes.Size())
	st.Push(start)

	for !st.IsEmpty() {
		node, err := st.Pop()
		if err != nil {
			// Must not reach here
			return false
		}

		//fmt.Printf("node:%v, edges=%+v\n", node, g.Edges(node))
		visitor.Visit(g, node)
		if node.IsEqual(target) {
			return true
		}
		dfs.visited[node] = struct{}{}

		for _, edge := range g.Edges(node) {
			if _, visited := dfs.visited[edge.end]; visited {
				continue
			}
			st.Push(edge.end)
			//fmt.Printf("Pushed: %+v\n", edge.end)
		}
	}

	return false
}
