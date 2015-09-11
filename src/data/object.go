package data

type Object struct {
	Value interface{}
}

func NewObjectInt(value int) Object {
	return Object{Value: value}
}

func (o Object) ToInt() int {
	return o.Value.(int)
}

func (o Object) ToIntDefault(defaultValue int) int {
	val, ok := o.Value.(int)
	if ok {
		return val
	}
	return defaultValue
}
