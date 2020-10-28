package queue

import (
	"sync"

	"github.com/jacygao/dsa/linkedlist"
)

type Queue struct {
	sync.Mutex
	l *linkedlist.LinkedList
}

func New(*Queue) {
	q: linkedlist.New(),
}

func (q *Queue) add(val interface{}){
	q.Lock()
	defer q.Unlock()
	q.l.Append(val)
}

func (q *Queue) remove() interface{} {
	q.Lock()
	defer q.Unlock()
	node := q.l.First
	copy := *node
	q.l.Delete(node)
	return copy
}

func (q *Queue) peek() interface{} {
	q.Lock()
	defer q.Unlock()
	return q.l.First
}

func (q *Queue) isEmpty() bool {
	q.Lock()
	defer q.Unlock()
	return q.l.First == nil
}