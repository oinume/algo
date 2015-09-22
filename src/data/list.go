package data

type List interface {
	Add(o Object) bool
	First() (Object, error)
	Size() int
	Iterator() Iterator
}
