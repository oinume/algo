package lru_cache

import (
	"io"
)

type LRUCache[K comparable, V any] interface {
	// Get returns value for `key` and true if found, or zero value and false if not found.
	Get(key K) (V, bool)
	// Put sets value with key.
	Put(key K, value V)
}

type Dumper interface {
	// Dump writes internal data of LRUCache to `w`.
	Dump(w io.Writer) error
}

type item[V any] struct {
	value V
	age   int
}
