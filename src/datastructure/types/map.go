package types

type Map interface {
	Put(key interface{}, value interface{}) (interface{}, error)
	Get(key interface{}) (interface{}, error)
	Size() int
	Remove(key interface{}) (interface{}, error)
}

type Hashable interface {
	HashCode() int
}
