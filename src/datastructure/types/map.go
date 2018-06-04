package types

type Map interface {
	Put(key Value, value Value) (Value, error)
	Get(key Value) (Value, error)
	Size() int
	Remove(key Value) (Value, error)
}

type Hashable interface {
	HashCode() int
}
