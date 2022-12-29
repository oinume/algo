package linked_list

import (
	"reflect"
	"testing"
)

func TestGenericLinkedList_Add(t *testing.T) {
	list := NewGenericLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(4)

	if got, want := list.Size(), 3; got != want {
		t.Errorf("unexpected size: got=%v, want=%v", got, want)
	}
}

func TestGenericLinkedList_Insert(t *testing.T) {
	list := NewGenericLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(4)
	list.Insert(1, 3)

	wantList := NewGenericLinkedList[int]()
	for i := 1; i <= 4; i++ {
		wantList.Add(i)
	}

	if got, want := list.Size(), 4; got != want {
		t.Errorf("unexpected size: got=%v, want=%v", got, want)
	}
	if !reflect.DeepEqual(list, wantList) {
		t.Errorf("unexpected linked list structure: got=%+v, want=%v", list, wantList)
	}
}

func TestGenericLinkedList_Set(t *testing.T) {
	list := NewGenericLinkedList[int]()
	for i := 1; i <= 3; i++ {
		list.Add(i)
	}
	old, err := list.Set(2, 300)
	if err != nil {
		t.Fatalf("Set() failed: %v", err)
	}
	if old != 3 {
		t.Fatalf("Set() returns unexpected value: got=%v, want=%v", old, 3)
	}

	var value interface{}
	for i := list.Iterator(); i.HasNext(); {
		v, err := i.Next()
		if err != nil {
			t.Fatalf("iterator Next() failed: %v", err)
		}
		value = v
	}
	if got, want := value, 300; got != want {
		t.Fatalf("last element is %v but must be %v", got, want)
	}
	//t.Logf("list = %v", spew.Sdump(list))

	oldFirst, err := list.Set(0, 100)
	if err != nil {
		t.Fatalf("Set() failed: %v", err)
	}
	if oldFirst != 1 {
		t.Fatalf("Set() returns unexpected value: got=%v, want=%v", oldFirst, 1)
	}

	if i := list.Iterator(); i.HasNext() {
		first, _ := i.Next()
		if got, want := first, 100; got != want {
			t.Errorf("first element is %v but must be %v", got, want)
		}
		second, _ := i.Next()
		if got, want := second, 2; got != want {
			t.Errorf("second element is %v but must be %v", got, want)
		}
	}
}

func TestGenericLinkedList_Remove(t *testing.T) {
	list := NewGenericLinkedList[int]()
	for i := 1; i <= 3; i++ {
		list.Add(i)
	}
	if !list.Remove(1) {
		t.Fatalf("Remove() must return true")
	}

	got, err := list.Iterator().Next()
	if err != nil {
		t.Fatalf("Iterator().Next() failed: %v", err)
	}
	if want := 2; got != want {
		t.Fatalf("Iterator().Next() returns unexpected value: got=%v, want=%v", got, want)
	}

	tests := []struct {
		value   int
		removed bool
	}{
		{2, true},
		{3, true},
		{100, false},
	}
	for _, test := range tests {
		if got, want := list.Remove(test.value), test.removed; got != want {
			t.Errorf("Remove() return unexpected reuslt: got=%v, want=%v", got, want)
		}
	}
	if list.Size() != 0 {
		t.Errorf("list must be empty but size is %v", list.Size())
	}
	if got, want := list.Iterator().HasNext(), false; got != want {
		t.Errorf("list iterator HasNext() must be false but %v", got)
	}
}

func TestGenericLinkedListIterator_Remove(t *testing.T) {
	list := NewGenericLinkedList[int]()
	for i := 1; i <= 3; i++ {
		list.Add(i)
	}

	iterator := list.Iterator()
	first, err := iterator.Next()
	if err != nil {
		t.Fatalf("iterator.Next() failed: %v", err)
	}

	if got, want := first, 1; got != want {
		t.Errorf("Add() failed: got=%v, want=%v", got, want)
	}
	for i := 1; i <= 3; i++ {
		removed, err := iterator.Remove()
		if err != nil {
			t.Fatalf("iterator.Remove() failed: %v", err)
		}
		if got, want := removed, i; got != want {
			t.Errorf("unexpected removed value: got=%v, want=%v", got, want)
		}
	}
}
