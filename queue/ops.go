package queue

type Operations interface {
	Add(interface{})
	Remove() interface{}
	Peek() interface{}
	IsEmpty() bool
}
