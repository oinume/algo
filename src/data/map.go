package data

type Map interface {
	Put(key Value, value Value)
	Get(key Value)
}
