package hash_table

import "errors"

var (
	ErrNotExists       = errors.New("cannot find data for the key")
	ErrKeyMustNotBeNil = errors.New("key must not be nil")
	ErrHashTableIsFull = errors.New("hash table is full")
)
