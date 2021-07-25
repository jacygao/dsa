package queue

import (
	"sync"

	"github.com/jacygao/dsa/linkedlist"
)

type Queue struct {
	sync.Mutex
	l *linkedlist.LinkedList
}

func New() *Queue {
	return &Queue{
		l: linkedlist.New(),
	}
}

func (q *Queue) Add(val interface{}) {
	q.Lock()
	defer q.Unlock()
	q.l.Append(val)
}

func (q *Queue) Remove() interface{} {
	q.Lock()
	defer q.Unlock()
	node := q.l.First
	copy := *node
	q.l.Delete(node)
	return copy.Value()
}

func (q *Queue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()
	return q.l.First.Value()
}

func (q *Queue) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()
	return q.l.First == nil
}
