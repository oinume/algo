package data

import "fmt"

type element struct {
	data Object
	next *element
}

type linkedList struct {
	first *element
}

func NewLinkedList() List {
	return &linkedList{first: &element{data: Object{}, next: nil}}
}

func (l *linkedList) Add(o Object) bool {
	lastElement := l.first
	for e := l.first.next; e != nil; e = e.next {
		fmt.Printf("e = %v\n", e.data)
		lastElement = e
	}
	fmt.Printf("lastElement = %v\n", lastElement.data)
	return false
}

func (l *linkedList) HasNext() bool {
	return false
}

func (l *linkedList) Next() (Object, error) {
	return Object{}, nil
}
