package btree

type Node interface {
	GetKey() int
	GetValue() interface{}
}

type NodeImpl struct{}

type Leaf struct {
	Key   int
	Value interface{}
}

func (l *Leaf) GetKey() int {
	return l.Key
}

func (l *Leaf) GetValue() interface{} {
	return l.Value
}

func NewLeaf(key int, value interface{}) Node {
	return &Leaf{
		Key:   key,
		Value: value,
	}
}
