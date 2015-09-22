package data

type List interface {
	Add(o Object) bool
	Size() int
	Iterator() Iterator
}
