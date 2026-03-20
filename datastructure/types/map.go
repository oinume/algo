package types

type Map[K comparable, V any] interface {
	Put(key K, value V) (V, error)
	Get(key K) (V, error)
	Size() int
	Remove(key K) (V, error)
}

type Hashable interface {
	HashCode() int
}
