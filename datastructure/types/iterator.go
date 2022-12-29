package types

type Iterator[T any] interface {
	Next() (T, error)
	HasNext() bool
	Remove() (T, error)
}
