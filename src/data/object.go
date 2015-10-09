package data

import (
	"fmt"
)

type Object struct {
	Value interface{}
}

func NewObjectInt(value int) Object {
	return Object{Value: value}
}

func (o Object) String() string {
	return fmt.Sprintf("Object{Value: %v}", o.Value)
}

func (o Object) ToInt() int {
	return o.Value.(int)
}

func (o Object) ToIntDefault(defaultValue int) int {
	value, ok := o.Value.(int)
	if ok {
		return value
	}
	return defaultValue
}

func (o Object) ToString() string {
	return o.Value.(string)
}

func (o Object) ToStringDefault(defaultValue string) string {
	value, ok := o.Value.(string)
	if ok {
		return value
	}
	return defaultValue
}

