package types

type Value interface {
	Get() interface{}
	Receive(v interface{}) error
	String() string
	Int() int
}