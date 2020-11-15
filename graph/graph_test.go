package graph

import (
	"fmt"
	"math/rand"
	"reflect"
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

func TestCycleExists(t *testing.T) {
	n := 4
	e := 6
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 0},
		{2, 3},
		{3, 3},
	}

	expected := true
	if res := cycleExists(n, e, edges); res != expected {
		t.Fatalf("results do not match! expected %v but got %v", expected, res)
	}
}

func TestCycleNotExists(t *testing.T) {
	n := 4
	e := 3
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 3},
	}

	expected := false
	if res := cycleExists(n, e, edges); res != expected {
		t.Fatalf("results do not match! expected %v but got %v", expected, res)
	}
}

func TestFindBridges(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{1, 5},
	}

	expected := [][]int{
		{3, 4},
		{0, 1},
	}
	ans := findBridges(n, edges)

	if !reflect.DeepEqual(ans, expected) {
		t.Fatalf("results do not match! expected %+v but got %+v", expected, ans)
	}
}

func TestFindBridgesWithNestedNodes(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{2, 3},
		{2, 4},
		{3, 4},
	}

	expected := [][]int{
		{0, 1},
	}
	ans := findBridges(n, edges)

	if !reflect.DeepEqual(ans, expected) {
		t.Fatalf("results do not match! expected %+v but got %+v", expected, ans)
	}
}

func TestFindBridgesWithCompositeNodes(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
		{4, 5},
	}

	expected := [][]int{
		{1, 3},
	}
	ans := findBridges(n, edges)

	if !reflect.DeepEqual(ans, expected) {
		t.Fatalf("results do not match! expected %+v but got %+v", expected, ans)
	}
}
