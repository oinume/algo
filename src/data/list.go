package data

type List interface {
	Add(o Object) bool
	HasNext() bool
	Next() (Object, error)
}
