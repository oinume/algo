package types

type Iterator interface {
	Next() (interface{}, error)
	HasNext() bool
	Remove() (interface{}, error)
}
