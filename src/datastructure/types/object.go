package types

// Object is an implementation of `Value`
//type Object struct {
//	Value interface{}
//}
//
//func (o *Object) Get() interface{} {
//	return o.Value
//}
//
//func (o *Object) Receive(v interface{}) error {
//	// TODO
//	return nil
//}
//
//func (o Object) Int() int {
//	return o.Value.(int)
//}
//
//func (o Object) IntDefault(def int) int {
//	value, ok := o.Value.(int)
//	if ok {
//		return value
//	}
//	return def
//}
//
//func (o *Object) String() string {
//	//return fmt.Sprintf("Object{Value: %v}", o.Value)
//	return fmt.Sprint(o.Value)
//}
//
//func (o Object) StringDefault(def string) string {
//	value, ok := o.Value.(string)
//	if ok {
//		return value
//	}
//	return def
//}
