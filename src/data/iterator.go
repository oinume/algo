package data

type Iterator interface {
	Next() (Object, error)
	HasNext() bool
	Remove() (Object, error)
}
