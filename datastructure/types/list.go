package types

import "fmt"

type List interface {
	// Add o as last into this list
	Add(v interface{})
	// Insert o after the position of index
	Insert(index int, v interface{})
	// Get size
	Size() int
	// Remove given v
	Remove(v interface{}) bool
	Set(index int, v interface{}) (interface{}, error)
	Iterator() Iterator
}

var (
	ErrorIndexOutOfRange = fmt.Errorf("index out of range")
)
