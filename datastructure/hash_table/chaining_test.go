package hash_table_test

import (
	"testing"

	"github.com/oinume/algo/datastructure/hash_table"
	"github.com/oinume/algo/testings"
)

//type hashable struct {
//	object types.Object
//}
//
//func (h *hashable) Get() interface{} {
//	return h.object
//}
//
//func (h *hashable) Receive(v interface{}) error {
//	return h.object.Receive(v)
//}
//
//func (h *hashable) String() string {
//	return h.object.String()
//}
//
//func (h *hashable) Int() int {
//	return h.object.Int()
//}

//// Always return same hash code
//func (h *hashable) HashCode() int {
//	return 1
//}

func TestHashTableChaining_Put(t *testing.T) {
	table := hash_table.NewChaining(10)
	_, err := table.Put(1, 1)
	if err != nil {
		t.Fatalf("Put: unexpected error: %v", err)
	}
	testings.AssertEqual(t, 1, table.Size(), "table.Size()")
}

func TestHashTableChaining_Put_Collision(t *testing.T) {
	table := hash_table.NewChaining(10)
	_, _ = table.Put("abc", "ABC")
	_, _ = table.Put("cba", "CBA")
	testings.AssertEqual(t, 2, table.Size(), "table.Size()")

	got, err := table.Get("cba")
	if err != nil {
		t.Fatalf("Get returns unexpected error: %v", err)
	}
	testings.AssertEqual(t, "CBA", got, "table.Get()")
}

func TestHashTableChaining_Put_Collision_Exists(t *testing.T) {
	table := hash_table.NewChaining(10)
	_, _ = table.Put("abc", "ABC")
	_, _ = table.Put("abc", "AABBCC")
	_, _ = table.Put("cba", "CBA")
	testings.AssertEqual(t, 2, table.Size(), "table.Size()")

	got, err := table.Get("abc")
	if err != nil {
		t.Fatalf("Get returns unexpected error: %v", err)
	}
	testings.AssertEqual(t, "AABBCC", got, "table.Get()")
}

func TestHashTableChaining_Get(t *testing.T) {
	table := hash_table.NewChaining(10)
	_, err := table.Put(1, 1)
	if err != nil {
		t.Fatalf("Put returns unexpected error: %v", err)
	}
	got, err := table.Get(1)
	if err != nil {
		t.Fatalf("Get returns unexpected error: %v", err)
	}
	testings.AssertEqual(t, 1, got, "table.Get()")
}

func TestHashTableChaining_Remove(t *testing.T) {
	table := hash_table.NewChaining(10)

	_, err := table.Put(1, 1)
	if err != nil {
		t.Fatalf("Put returns unexpected error: %v", err)
	}
	removed, err := table.Remove(1)
	if err != nil {
		t.Fatalf("Remove returns unexpected error: %v", err)
	}
	testings.AssertEqual(t, 1, removed, "table.Remove()")
	testings.AssertEqual(t, 0, table.Size(), "table must be empty")

	if _, err := table.Remove(100); err == nil {
		t.Fatalf("Remove must return err but nil")
	}
}
