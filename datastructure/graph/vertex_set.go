package graph

type vertexSet struct {
	values map[*Vertex]struct{}
}

func NewVertexSet(size int) *vertexSet {
	return &vertexSet{
		values: make(map[*Vertex]struct{}, size),
	}
}
