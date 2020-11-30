package graph

import (
	"github.com/jacygao/dsa/queue"
)

type Node struct {
	val      int
	children []*Node
	visited  bool
}

func (n *Node) AppendChild(node *Node) *Node {
	n.children = append(n.children, node)
	return node
}

func DFS(node *Node, visit func(n *Node)) {
	dfs(node, visit)
}

func dfs(node *Node, visit func(n *Node)) {
	visit(node)
	node.visited = true
	for _, n := range node.children {
		if !n.visited {
			dfs(n, visit)
		}
	}
}

func BFS(node *Node, visit func(n *Node)) {
	visit(node)
	node.visited = true
	bfs(node, visit)
}

func bfs(node *Node, visit func(n *Node)) {
	q := queue.New()
	node.visited = true
	q.Add(node)
	for !q.IsEmpty() {
		visit(q.Remove().(*Node))
		for _, n := range node.children {
			if !n.visited {
				n.visited = true
				q.Add(n)
			}
		}
	}
}
