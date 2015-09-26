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
			data: &Object{Value: 0},
			next: nil,
		},
	}
}

func (l *linkedList) Add(o Object) bool {
	lastElement := l.head
	for e := l.head.next; e != nil; e = e.next {
		lastElement = e
	}
	lastElement.next = &element{data: &o, next: nil}
	return true
}

func (l *linkedList) Size() int {
	var size int = 0
	for e := l.head.next; e != nil; e = e.next {
		size++
	}
	return size
}

func (l *linkedList) First() (Object, error) {
	if l.head.next != nil {
		return *(l.head.next.data), nil
	}
	return Object{0}, fmt.Errorf("Empty list.")
}

func (l *linkedList) Iterator() Iterator {
	return &linkedListIterator{cursor: l.head}
}

type linkedListIterator struct {
	cursor *element
}

func (i *linkedListIterator) Next() (Object, error) {
	if i.HasNext() {
		data := i.cursor.next.data
		i.cursor = i.cursor.next
		return *data, nil
	} else {
		return Object{0}, fmt.Errorf("No next element.")
	}
}

func (i *linkedListIterator) HasNext() bool {
	return i.cursor.next != nil
}

func (i *linkedListIterator) Remove() (Object, error) {
	data := i.cursor.data
	if data == nil {
		return Object{0}, fmt.Errorf("No current object")
	}
	if i.HasNext() {
		i.cursor = i.cursor.next
	}
	return *data, nil
}
