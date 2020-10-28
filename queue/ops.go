package queue

type Operations interface {
	add(interface{})
	remove() interface{}
	peek() interface{}
	isEmpty() bool
}
