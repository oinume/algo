package data

import (
	"fmt"
)

type element struct {
	data Value
	next *element
}

func (e *element) String() string {
	if e.next != nil {
		return fmt.Sprintf("{data: %v, next: %v}", e.data, e.next.data)
	} else {
		return fmt.Sprintf("{data: %v, next: nil}", e.data)
	}
}

type linkedList struct {
	head *element
}

func NewLinkedList() List {
	return &linkedList{
		head: &element{
			data: nil,
			next: nil,
		},
	}
}

func (l *linkedList) Add(v Value) {
	last := l.head
	for e := l.head.next; e != nil; e = e.next {
		last = e
	}
	last.next = &element{data: v, next: nil}
}

func (l *linkedList) Insert(index int, v Value) {
	current := 0
	for e := l.head.next; e != nil; e = e.next {
		if current == index {
			next := e.next
			e.next = &element{data: v, next: next}
			break
		}
		current++
	}
}

func (l *linkedList) Size() int {
	// TODO: This code is too slow
	var size int = 0
	for e := l.head.next; e != nil; e = e.next {
		size++
	}
	return size
}

func (l *linkedList) Set(index int, v Value) (Value, error) {
	if index >= l.Size() {
		return nil, ErrorIndexOutOfRange
	}

	i := 0
	current, prev := l.head, l.head
	for e := l.head.next; e != nil; e = e.next {
		prev = current
		current = e
		if i == index {
			break
		}
	}
	oldElement := current
	oldElement.next = nil
	newElement := &element{data: v, next: oldElement.next}
	prev.next = newElement

	return oldElement.data, nil
}

func (l *linkedList) Remove(v Value) bool {
	prev := l.head
	var target *element = nil
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
	target.data = nil
	target.next = nil
	//fmt.Printf("target = %+v, prev = %+v\n", target, prev)

	return true
}

func (l *linkedList) First() (Value, error) {
	if l.head.next != nil {
		return l.head.next.data, nil
	}
	return nil, fmt.Errorf("Empty list.")
}

func (l *linkedList) Iterator() Iterator {
	return &linkedListIterator{cursor: l.head}
}

//
// Iterator
//
type linkedListIterator struct {
	cursor *element
}

func (i *linkedListIterator) Next() (Value, error) {
	if i.HasNext() {
		data := i.cursor.next.data
		i.cursor = i.cursor.next
		return data, nil
	} else {
		return nil, fmt.Errorf("No next element.")
	}
}

func (i *linkedListIterator) HasNext() bool {
	return i.cursor.next != nil
}

func (i *linkedListIterator) Remove() (Value, error) {
	data := i.cursor.data
	if data == nil {
		return nil, fmt.Errorf("No current object")
	}
	if i.HasNext() {
		i.cursor = i.cursor.next
	}
	return data, nil
}
