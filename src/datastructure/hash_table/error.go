package hash_table

import "errors"

var (
	ErrKeyNotExists    = errors.New("key not exists")
	ErrKeyMustNotBeNil = errors.New("key must not be nil")
	ErrHashTableIsFull = errors.New("hash table is full")
)
