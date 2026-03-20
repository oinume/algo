package graph

type nodeSet struct {
	values map[*Node]struct{}
}

func NewNodeSet(size int) *nodeSet {
	return &nodeSet{
		values: make(map[*Node]struct{}, size),
	}
}

func (vs *nodeSet) Add(v *Node) {
	vs.values[v] = struct{}{}
}

func (vs *nodeSet) Size() int {
	return len(vs.values)
}

func (vs *nodeSet) Values() []*Node {
	ret := make([]*Node, vs.Size())
	for v := range vs.values {
		ret = append(ret, v)
	}
	return ret
}
