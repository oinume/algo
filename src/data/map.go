package data

type Map interface {
	Put(key Value, value Value) Value
	Get(key Value) (Value, error)
}
