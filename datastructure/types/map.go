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

type GenericMap[K Hashable, V] interface {
	Put(key K, value V) (V, error)
	Get(key K) (V, error)
	Size() int
	Remove(key K) (V, error)
}
