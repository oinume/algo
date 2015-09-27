package data

import "fmt"

type List interface {
	Add(o Object) bool
	Size() int
	//Remove(o Object) bool
	Set(index int, o Object) (Object, error)
	Iterator() Iterator
}

var (
	ErrorIndexOutOfRange error = fmt.Errorf("Index out of range")
)
