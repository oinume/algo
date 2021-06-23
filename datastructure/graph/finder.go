package graph

import (
	"github.com/oinume/algo/datastructure/stack"
)

type Finder interface {
	Find(g *Graph, start *Vertex, target *Vertex, visitor Visitor) bool
}

// dfsRecursiveFinder is depth first search finder
type dfsRecursiveFinder struct {
	visited map[*Vertex]struct{}
}

func NewDFSRecursiveFinder() Finder {
	return &dfsRecursiveFinder{
		visited: make(map[*Vertex]struct{}, 100),
	}
}

func (dfs *dfsRecursiveFinder) Find(g *Graph, start *Vertex, target *Vertex, visitor Visitor) bool {
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
	visited map[*Vertex]struct{}
}

func NewDFSLoopFinder() Finder {
	return &dfsLoopFinder{
		visited: make(map[*Vertex]struct{}, 100),
	}
}

func (dfs *dfsLoopFinder) Find(g *Graph, start *Vertex, target *Vertex, visitor Visitor) bool {
	st := stack.New(g.vertices.Size())
	st.Push(start)

	for !st.IsEmpty() {
		v, err := st.Pop()
		if err != nil {
			// Must not reach here
			return false
		}

		vertex := v.(*Vertex)
		//fmt.Printf("vertex:%v, edges=%+v\n", vertex, g.Edges(vertex))
		visitor.Visit(g, vertex)
		if vertex.IsEqual(target) {
			return true
		}
		dfs.visited[vertex] = struct{}{}

		for _, edge := range g.Edges(vertex) {
			if _, visited := dfs.visited[edge.end]; visited {
				continue
			}
			st.Push(edge.end)
			//fmt.Printf("Pushed: %+v\n", edge.end)
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
