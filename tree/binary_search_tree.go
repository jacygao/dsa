package tree

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

func (t *BinarySearchTree) TravelInOrder(visit func(node *Node)) {
	t.travelInOrder(t.root, visit)
}

func (t *BinarySearchTree) TravelPreOrder(visit func(node *Node)) {
	t.travelPreOrder(t.root, visit)
}

func (t *BinarySearchTree) TravelPostOrder(visit func(node *Node)) {
	t.travelPostOrder(t.root, visit)
}

func (t *BinarySearchTree) travelInOrder(node *Node, visit func(node *Node)) {
	if node != nil {
		t.travelInOrder(node.left, visit)
		visit(node)
		t.travelInOrder(node.right, visit)
	}
}

func (t *BinarySearchTree) travelPreOrder(node *Node, visit func(node *Node)) {
	if node != nil {
		visit(node)
		t.travelInOrder(node.left, visit)
		t.travelInOrder(node.right, visit)
	}
}

func (t *BinarySearchTree) travelPostOrder(node *Node, visit func(node *Node)) {
	if node != nil {
		t.travelInOrder(node.left, visit)
		t.travelInOrder(node.right, visit)
		visit(node)
	}
}

func NewBinarySearchTreeWithMinimumDepth(vals []int) *BinarySearchTree {
	tree := &BinarySearchTree{}
	length := len(vals)
	insertVal(tree, 0, length-1, vals)
	return tree
}

func insertVal(tree *BinarySearchTree, start, end int, vals []int) {
	if start <= end {
		index := int((start + end) / 2)
		tree.Add(vals[index])
		// left
		insertVal(tree, start, index-1, vals)
		// right
		insertVal(tree, index+1, end, vals)
	}
}

func (t *BinarySearchTree) MaxDepth() int {
	return t.walk(t.root)
}

func (t *BinarySearchTree) walk(node *Node) int {
	if node != nil {
		lDep := t.walk(node.left)
		rDep := t.walk(node.right)

		if lDep > rDep {
			return lDep + 1
		}
		return rDep + 1
	}
	return 0
}

type LNode struct {
	prev LNode
	next LNode
	val  int
}

func NewTreeToLinkedLists() {

}

fucn walkTree()