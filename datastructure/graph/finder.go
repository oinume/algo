package graph

type Finder interface {
	Find(g *Graph, start *Vertex, target *Vertex, visitor Visitor) bool
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

func (dfs *dfsFinder) Find(g *Graph, start *Vertex, target *Vertex, visitor Visitor) bool {
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

type Visitor interface {
	Visit(g *Graph, v *Vertex)
}

type nopVisitor struct{}

func (nv *nopVisitor) Visit(g *Graph, v *Vertex) {}

func (nv *nopVisitor) Visited() []*Vertex {
	return nil
}

func NewListVisitor() Visitor {
	return &listVisitor{
		list: make([]*Vertex, 0, 100),
	}
}

type listVisitor struct {
	list []*Vertex
}

func (lv *listVisitor) Visit(g *Graph, v *Vertex) {
	lv.list = append(lv.list, v)
}
