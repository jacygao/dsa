package graph

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDFS(t *testing.T) {
	vals := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	nodes := []*Node{}
	for _, val := range vals {
		nodes = append(nodes, &Node{
			val: val,
		})
	}

	root := nodes[0]
	for i := 0; i < 10; i++ {
		nodesTotal := len(vals)
		index := rand.Intn(nodesTotal)
		for index == 0 || root == nodes[index-1] {
			index = rand.Intn(nodesTotal)
		}
		root = root.AppendChild(nodes[index-1])
	}

	visit := func(n *Node) {
		fmt.Println(n.val)
	}
	DFS(nodes[0], visit)
}

func TestBFS(t *testing.T) {
	vals := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	nodes := []*Node{}
	for _, val := range vals {
		nodes = append(nodes, &Node{
			val: val,
		})
	}

	root := nodes[0]
	for i := 0; i < 10; i++ {
		nodesTotal := len(vals)
		index := rand.Intn(nodesTotal)
		for index == 0 || root == nodes[index-1] {
			index = rand.Intn(nodesTotal)
		}
		fmt.Printf("%d ", nodes[index-1].val)
		root = root.AppendChild(nodes[index-1])
	}
	fmt.Printf("\n result \n")
	visit := func(n *Node) {
		fmt.Printf("%d ", n.val)
	}
	BFS(nodes[0], visit)
}
