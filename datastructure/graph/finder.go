package graph

import (
	"github.com/oinume/algo/datastructure/stack"
)

type Finder[T comparable] interface {
	Find(g *Graph[T], start *Node[T], target *Node[T], visitor Visitor[T]) bool
}

// dfsRecursiveFinder is depth first search finder
type dfsRecursiveFinder[T comparable] struct {
	visited map[*Node[T]]struct{}
}

func NewDFSRecursiveFinder[T comparable]() Finder[T] {
	return &dfsRecursiveFinder[T]{
		visited: make(map[*Node[T]]struct{}, 100),
	}
}

func (dfs *dfsRecursiveFinder[T]) Find(g *Graph[T], start *Node[T], target *Node[T], visitor Visitor[T]) bool {
	visitor.Visit(g, start)

	if start.IsEqual(target) {
		return true
	}
	if _, visited := dfs.visited[start]; visited {
		return false
	}

	dfs.visited[start] = struct{}{}
	edges := g.Edges(start)
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

type dfsLoopFinder[T comparable] struct {
	visited map[*Node[T]]struct{}
}

func NewDFSLoopFinder[T comparable]() Finder[T] {
	return &dfsLoopFinder[T]{
		visited: make(map[*Node[T]]struct{}, 100),
	}
}

func (dfs *dfsLoopFinder[T]) Find(g *Graph[T], start *Node[T], target *Node[T], visitor Visitor[T]) bool {
	st := stack.New[*Node[T]](g.nodes.Size())
	st.Push(start)

	for !st.IsEmpty() {
		node, err := st.Pop()
		if err != nil {
			// Must not reach here
			return false
		}

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
		}
	}

	return false
}

// bfsRecursiveFinder is breadth first search finder
type bfsRecursiveFinder struct {
	visited map[*Vertex]struct{}
}

func NewBFSRecursiveFinder() Finder {
	return &bfsRecursiveFinder{
		visited: make(map[*Vertex]struct{}, 100),
	}
}

func (dfs *bfsRecursiveFinder) Find(g *Graph, start *Vertex, target *Vertex, visitor Visitor) bool {
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
