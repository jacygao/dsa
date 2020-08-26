package linkedlist

// LinkedList wraps a simple doubly linked list.
type LinkedList struct {
	First *Node
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
		return
	}

	last := l.First.Next()
	if last == nil {
		node.prev = l.First
		l.First.next = node
		return
	}

	for last.Next() != nil {
		last = last.Next()
	}

	last.next = node
	node.prev = last
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
		n.prev.next = nil
		return
	}

	n.prev.next = n.Next()
	n.next.prev = n.Prev()
}

// Last returns the last node in a linked list.
func (l *LinkedList) Last() *Node {
	last := l.First
	if last == nil {
		return nil
	}

	for last.Next() != nil {
		last = last.Next()
	}

	return last
}

// ShiftToEnd shifts a node to the end of a linked list.
func (l *LinkedList) ShiftToEnd(n *Node) {
	if n == nil {
		return
	}

	next := n.Next()
	if next == nil {
		// n is alread the last node so nothing needs to be done
		return
	}

	last := l.Last()
	prev := n.Prev()

	if prev == nil {
		// n is the first node on the list
		l.First = next
		next.prev = nil
	} else {
		prev.next = next
		next.prev = prev
	}

	last.next = n
	n.next = nil
	n.prev = last
}
