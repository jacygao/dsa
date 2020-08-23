package linkedlist

// LinkedList wraps a simple doubly linked list.
type LinkedList struct {
	First *Node
	Last  *Node
}

// Node defines a node in a linked list
type Node struct {
	val  interface{}
	next *Node
	prev *Node
}

// Value returns the value of a node.
func (n *Node) Value() interface{} {
	return n.val
}

// Next returns the next node of a node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous node of a node.
func (n *Node) Prev() *Node {
	return n.prev
}

// New returns an empty doubly linked list.
func New() *LinkedList {
	return &LinkedList{}
}

// Append appends a new node with given value to the end of the linked list.
func (l *LinkedList) Append(data interface{}) {
	node := &Node{}
	node.val = data

	if l.First == nil {
		l.First = node
		l.Last = node
		return
	}

	last := l.Last
	last.next = node
	node.prev = last
	l.Last = node
}

// Delete removes a node from the linked list
func (l *LinkedList) Delete(n *Node) {
	// Deleting first node
	if n.Prev() == nil {
		l.First = n.Next()
		n.next.prev = nil
		return
	}
	// Deleting last node
	if n.Next() == nil {
		l.Last = n.Prev()
		n.prev.next = nil
		return
	}

	n.prev.next = n.Next()
	n.next.prev = n.Prev()
}
