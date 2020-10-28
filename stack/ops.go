package stack

// Operations defines the operations used by a stack.
type Operations interface {
	Push(data interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
}
