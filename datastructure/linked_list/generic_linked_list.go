package linked_list

import (
	"fmt"

	"github.com/oinume/algo/datastructure/types"
)

type genericElement[T any] struct {
	data T
	next *genericElement[T]
}

func (e *genericElement[T]) String() string {
	if e.next != nil {
		return fmt.Sprintf("{data: %v, next: %v}", e.data, e.next.data)
	} else {
		return fmt.Sprintf("{data: %v, next: nil}", e.data)
	}
}

type genericLinkedList[T any] struct {
	head *genericElement[T]
}

// var _ types.GenericList = (*genericLinkedList)(nil)

func NewGenericLinkedList[T any]() types.GenericList[T] {
	var data T
	return &genericLinkedList[T]{
		head: &genericElement[T]{
			data: data,
			next: nil,
		},
	}
}

func (l *genericLinkedList[T]) Add(v T) {
	last := l.head
	for e := l.head.next; e != nil; e = e.next {
		last = e
	}
	last.next = &genericElement[T]{data: v, next: nil}
}

func (l *genericLinkedList[T]) Insert(index int, v T) {
	current := 0
	for e := l.head.next; e != nil; e = e.next {
		if current == index {
			next := e.next
			e.next = &genericElement[T]{data: v, next: next}
			break
		}
		current++
	}
}

func (l *genericLinkedList[T]) Size() int {
	// TODO: This code is too slow
	size := 0
	for e := l.head.next; e != nil; e = e.next {
		size++
	}
	return size
}

func (l *genericLinkedList[T]) Set(index int, v T) (T, error) {
	if index >= l.Size() {
		var empty T
		return empty, types.ErrorIndexOutOfRange
	}

	i := 0
	current, prev := l.head, l.head
	for e := l.head.next; e != nil; e = e.next {
		prev = current
		current = e
		if i == index {
			break
		}
		i++
	}

	oldValue := current.data
	next := current.next
	current.next = nil
	newElement := &genericElement[T]{data: v, next: next}
	prev.next = newElement

	return oldValue, nil
}

func (l *genericLinkedList[T]) Remove(v T) bool {
	prev := l.head
	var target *genericElement[T] = nil
	for e := l.head; e != nil; e = e.next {
		if fmt.Sprint(e.data) == fmt.Sprint(v) { // TODO: Equals
			target = e
			break
		}
		prev = e
		//fmt.Printf("e = %+v, prev = %+v\n", e, prev)
	}
	if target == nil {
		return false
	}
	prev.next = target.next
	var empty T
	target.data = empty
	target.next = nil
	//fmt.Printf("target = %+v, prev = %+v\n", target, prev)

	return true
}

func (l *genericLinkedList[T]) First() (T, error) {
	if l.head.next != nil {
		return l.head.next.data, nil
	}
	var empty T
	return empty, fmt.Errorf("empty list")
}

func (l *genericLinkedList[T]) Iterator() types.GenericIterator[T] {
	return &genericLinkedListIterator[T]{cursor: l.head}
}

//
// Iterator
//
type genericLinkedListIterator[T any] struct {
	cursor *genericElement[T]
}

func (i *genericLinkedListIterator[T]) Next() (T, error) {
	if i.HasNext() {
		data := i.cursor.next.data
		i.cursor = i.cursor.next
		return data, nil
	} else {
		var empty T
		return empty, fmt.Errorf("no next element")
	}
}

func (i *genericLinkedListIterator[T]) HasNext() bool {
	return i.cursor.next != nil
}

func (i *genericLinkedListIterator[T]) Remove() (T, error) {
	data := i.cursor.data
	// TODO: Fix invalid operation: data == nil (mismatched types T and untyped nil)
	//if data == nil {
	//	var empty T
	//	return empty, fmt.Errorf("no current object")
	//}
	if i.HasNext() {
		i.cursor = i.cursor.next
	}
	return data, nil
}
