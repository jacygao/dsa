package stack

import (
	"sync"

	"github.com/jacygao/linkedlist/linkedlist"
)

// Stack implements a thread safe stack backed by a linked list.
type Stack struct {
	sync.Mutex
	list *linkedlist.LinkedList
}

// New returns a new stack.
func New() *Stack {
	return &Stack{
		list: linkedlist.New(),
	}
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(data interface{}) {
	s.Lock()
	s.list.Prepend()
	s.Unlock()
}

// Pop returns a value on top of the stack.
func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()
	node := s.list.First
	copy := *node
	s.list.Delete(node)
	return copy
}

// Peek returns the top value of the stack.
func (s *Stack) Peek() interface{} {
	s.Lock()
	defer s.Unlock()
	return s.list.First
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()
	return s.list.First == nil
}
