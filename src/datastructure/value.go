package datastructure

import (
	"fmt"
)

type Value interface {
	Get() interface{}
	Receive(v interface{}) error
	ToString() string
	ToInt() int
}

// Object is an implementation of `Value`
type Object struct {
	Value interface{}
}

func (o *Object) String() string {
	return fmt.Sprintf("Object{Value: %v}", o.Value)
}

func (o *Object) Get() interface{} {
	return o.Value
}

func (o *Object) Receive(v interface{}) error {
	// TODO
	return nil
}

func (o Object) ToInt() int {
	return o.Value.(int)
}

func (o Object) ToIntDefault(def int) int {
	value, ok := o.Value.(int)
	if ok {
		return value
	}
	return def
}

func (o Object) ToString() string {
	// TODO: FIX: panic: interface conversion: interface is int, not string
	return o.Value.(string)
}

func (o Object) ToStringDefault(def string) string {
	value, ok := o.Value.(string)
	if ok {
		return value
	}
	return def
}

