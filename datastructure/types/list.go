package types

import "fmt"

type List interface {
	// Add adds v as last into this list
	Add(v interface{})
	// Insert inserts o after the position of index
	Insert(index int, v interface{})
	// Size returns size of this list
	Size() int
	// Remove removes given v
	Remove(v interface{}) bool
	Set(index int, v interface{}) (interface{}, error)
	Iterator() Iterator
}

type GenericList[T any] interface {
	Add(v T)
	Insert(index int, v T)
	Size() int
	Remove(v T) bool
	Set(index int, v T) (T, error)
	Iterator() GenericIterator[T]
}

var (
	ErrorIndexOutOfRange = fmt.Errorf("index out of range")
)
