package data

import (
	"testing"

	_ "github.com/stretchr/testify/assert"
)

func TestHashMapPut(t *testing.T) {
	hashMap := NewHashMap(10)
	hashMap.Put(&Object{1}, &Object{1})
}

