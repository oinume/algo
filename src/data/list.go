package data

import "fmt"

type List interface {
	// Add o as last into this list
	Add(v Value)
	// Insert o after the position of index
	Insert(index int, v Value)
	// Get size
	Size() int
	//Remove(o Object) bool
	Set(index int, v Value) (Value, error)
	Iterator() Iterator
}

var (
	ErrorIndexOutOfRange error = fmt.Errorf("Index out of range")
)
