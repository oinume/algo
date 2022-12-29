package types

import "fmt"

type List[T any] interface {
	Add(v T)
	Insert(index int, v T)
	Size() int
	Remove(v T) bool
	Set(index int, v T) (T, error)
	Iterator() Iterator[T]
}

var (
	ErrorIndexOutOfRange = fmt.Errorf("index out of range")
)
