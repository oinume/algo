package lru_cache

import (
	"io"
)

type LRUCache interface {
	// Get returns value for `key`. Returns -1 if not found
	Get(key int) int
	// Put sets value with key
	Put(key, value int)
}

type Dumper interface {
	// Dump writes internal data of LRUCache to `w`.
	Dump(w io.Writer) error
}

type item struct {
	value int
	age   int
}
