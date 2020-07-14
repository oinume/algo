package hash_table_test

import (
	"fmt"
	"testing"

	"github.com/oinume/algo/datastructure/hash_table"
	"github.com/oinume/algo/testings"
)

func TestOpenAddressing_Put(t *testing.T) {
	hashTable := hash_table.NewOpenAddressing()

	tests := map[string]struct {
		key        interface{}
		value      interface{}
		wantReturn interface{}
	}{
		"new key 1":      {key: 1, value: 10, wantReturn: nil},
		"new key 2":      {key: 2, value: 20, wantReturn: nil},
		"existing key 2": {key: 2, value: 30, wantReturn: 20},
	}

	for name, tc := range tests {
		ret, err := hashTable.Put(tc.key, tc.value)
		if err != nil {
			t.Fatalf("Put returns unexpected error: %v", err)
		}
		testings.AssertEqual(t, tc.wantReturn, ret, fmt.Sprintf("%v: Put returns unexpected value", name))
	}
	testings.AssertEqual(t, 2, hashTable.Size(), "Size")
}

func TestOpenAddressing_Put_Rehash(t *testing.T) {
	hashTable := hash_table.NewOpenAddressingWithMaxSize(3)

	tests := map[string]struct {
		key   interface{}
		value interface{}
	}{
		"abc": {key: "abc", value: "ABC"},
		"cba": {key: "cba", value: "CBA"},
	}

	for name, tt := range tests {
		got, err := hashTable.Put(tt.key, tt.value)
		if err != nil {
			t.Fatalf("%v: Put returns unexpected error: %v", name, err)
		}
		if got != nil {
			t.Fatalf("%v: Put must return nil but %v", name, got)
		}

		result, err := hashTable.Get(tt.key)
		if err != nil {
			t.Fatalf("Get returns unexpected error: %v", err)
		}
		testings.AssertEqual(t, tt.value, result, fmt.Sprintf("%v: Get returns unexpected value", name))
	}
}

func TestOpenAddressing_Get(t *testing.T) {
	hashTable := hash_table.NewOpenAddressing()

	tests := []struct {
		key   interface{}
		value interface{}
	}{
		{key: 1, value: 10},
		{key: 2, value: 20},
	}

	for _, tt := range tests {
		_, err := hashTable.Put(tt.key, tt.value)
		if err != nil {
			t.Fatalf("Put returns unexpected error: %v", err)
		}
		got, err := hashTable.Get(tt.key)
		if err != nil {
			t.Fatalf("Get returns unexpected error: %v", err)
		}
		testings.AssertEqual(t, tt.value, got, "Get")
	}
}

func TestOpenAddressing_Remove(t *testing.T) {
	hashTable := hash_table.NewOpenAddressing()

	tests := []struct {
		key    interface{}
		value  interface{}
		remove bool
	}{
		{key: "abc", value: "ABC", remove: true},
		{key: "cba", value: "CBA", remove: false},
	}

	for _, tt := range tests {
		_, err := hashTable.Put(tt.key, tt.value)
		if err != nil {
			t.Fatalf("Put returns unexpected error: %v", err)
		}
	}
	size := hashTable.Size()
	testings.AssertEqual(t, len(tests), size, "Size")

	for _, tt := range tests {
		if tt.remove {
			removed, err := hashTable.Remove(tt.key)
			if err != nil {
				t.Fatalf("Remove returns unexpected error: %v", err)
			}
			testings.AssertEqual(t, tt.value, removed, "Removed value")
			size--
		}
	}
	testings.AssertEqual(t, size, hashTable.Size(), "Size() must be decremented by removal")
}
