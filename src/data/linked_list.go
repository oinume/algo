package data

import (
	"fmt"
)

type element struct {
	data *Object
	next *element
}

type linkedList struct {
	head *element
}

func NewLinkedList() List {
	return &linkedList{
		head: &element{
			data: &Object{Value: -1},
			next: nil,
		},
	}
}

func (l *linkedList) Add(o *Object) {
	last := l.head
	for e := l.head.next; e != nil; e = e.next {
		last = e
	}
	last.next = &element{data: o, next: nil}
}

func (l *linkedList) Insert(index int, o *Object) {
	current := 0
	for e := l.head.next; e != nil; e = e.next {
		if current == index {
			next := e.next
			e.next = &element{data: o, next: next}
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

func (l *linkedList) Set(index int, o *Object) (*Object, error) {
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
	newElement := &element{data: o, next: oldElement.next}
	prev.next = newElement

	return oldElement.data, nil
}

func (l *linkedList) First() (*Object, error) {
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

func (i *linkedListIterator) Next() (*Object, error) {
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

func (i *linkedListIterator) Remove() (*Object, error) {
	data := i.cursor.data
	if data == nil {
		return nil, fmt.Errorf("No current object")
	}
	if i.HasNext() {
		i.cursor = i.cursor.next
	}
	return data, nil
}
