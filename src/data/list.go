package data

import "fmt"

type List interface {
	// Add o as last into this list
	Add(o *Object)
	// Insert o after the position of index
	Insert(index int, o *Object)
	// Get size
	Size() int
	//Remove(o Object) bool
	Set(index int, o *Object) (*Object, error)
	Iterator() Iterator
}

var (
	ErrorIndexOutOfRange error = fmt.Errorf("Index out of range")
)
