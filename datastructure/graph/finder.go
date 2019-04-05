package graph

type Finder interface {
	Find(g *Graph, start *Vertex, target *Vertex) bool
}

// dfsFinder is depth first search finder
type dfsFinder struct {
	visited map[*Vertex]struct{}
}

func NewDFSFinder() Finder {
	return &dfsFinder{
		visited: make(map[*Vertex]struct{}, 100),
	}
}

func (dfs *dfsFinder) Find(g *Graph, start *Vertex, target *Vertex) bool {
	if start.IsEqual(target) {
		return true
	}
	if _, ok := dfs.visited[start]; ok {
		return false
	}

	edges := g.Edges(start)
	for _, edge := range edges {
		if result := dfs.Find(g, edge.end, target); result {
			return result
		}
	}

	return false
}
