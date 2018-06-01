package types

type Iterator interface {
	Next() (Value, error)
	HasNext() bool
	Remove() (Value, error)
}
