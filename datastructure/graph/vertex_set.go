package graph

type vertexSet struct {
	values map[*Vertex]struct{}
}

func NewVertexSet(size int) *vertexSet {
	return &vertexSet{
		values: make(map[*Vertex]struct{}, size),
	}
}

func (vs *vertexSet) Add(v *Vertex) {
	vs.values[v] = struct{}{}
}

func (vs *vertexSet) Size() int {
	return len(vs.values)
}

func (vs *vertexSet) Values() []*Vertex {
	ret := make([]*Vertex, vs.Size())
	for v := range vs.values {
		ret = append(ret, v)
	}
	return ret
}
