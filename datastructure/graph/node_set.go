package graph

type nodeSet[T comparable] struct {
	values map[*Node[T]]struct{}
}

func newNodeSet[T comparable](size int) *nodeSet[T] {
	return &nodeSet[T]{
		values: make(map[*Node[T]]struct{}, size),
	}
}

func (vs *nodeSet[T]) Add(v *Node[T]) {
	vs.values[v] = struct{}{}
}

func (vs *nodeSet[T]) Size() int {
	return len(vs.values)
}

func (vs *nodeSet[T]) Values() []*Node[T] {
	ret := make([]*Node[T], 0, vs.Size())
	for v := range vs.values {
		ret = append(ret, v)
	}
	return ret
}
