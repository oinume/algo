package types

type Iterator interface {
	Next() (interface{}, error)
	HasNext() bool
	Remove() (interface{}, error)
}

type GenericIterator[T any] interface {
	Next() (T, error)
	HasNext() bool
	Remove() (T, error)
}
