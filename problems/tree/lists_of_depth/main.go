package main

import "fmt"

type Linkedlist struct {
	root *LNode
}

type LNode struct {
	prev *LNode
	next *LNode
	val  int
}

func (l *Linkedlist) add(val int) {
	if l.root == nil {
		l.root = &LNode{
			prev: nil,
			next: nil,
			val:  val,
		}
		return
	}
	n := l.root
	for n.next != nil {
		n = n.next
	}
	n.next = &LNode{
		prev: n,
		next: nil,
		val:  val,
	}
}

type Node struct {
	left  *Node
	right *Node
	Value int
}

type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) Add(val int) {
	t.root = t.add(t.root, val)
}

func (t *BinarySearchTree) add(root *Node, val int) *Node {
	if root == nil {
		return &Node{
			Value: val,
		}
	}
	// to the left
	if root.Value > val {
		root.left = t.add(root.left, val)
	}
	// to the right
	if root.Value < val {
		root.right = t.add(root.right, val)
	}
	return root
}

func (t *BinarySearchTree) walk(depth int, n *Node, visit func(depth int, n *Node)) {
	if n != nil {
		depth++
		t.walk(depth, n.left, visit)
		visit(depth, n)
		t.walk(depth, n.right, visit)
	}
}

func main() {
	tree := &BinarySearchTree{}
	nums := []int{
		5, 3, 12, 7, 2, 4, 8, 6, 10, 1,
	}

	for _, num := range nums {
		tree.Add(num)
	}

	lists := map[int]*Linkedlist{}
	visit := func(depth int, node *Node) {
		if _, ok := lists[depth]; !ok {
			lists[depth] = &Linkedlist{}
		}
		lists[depth].add(node.Value)
	}

	tree.walk(0, tree.root, visit)

	fmt.Printf("the length of the list is: %d \n", len(lists))
	for i, l := range lists {
		fmt.Printf("linkedlist %d \n", i)
		n := l.root
		for n != nil {
			fmt.Printf("value: %d \n", n.val)
			n = n.next
		}
	}
}
