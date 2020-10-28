package tree

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tree := &BinarySearchTree{}
	nums := []int{
		5, 3, 12, 7, 2, 4, 8, 6, 10, 1,
	}

	for _, num := range nums {
		tree.Add(num)
	}

	visit := func(node *Node) {
		fmt.Println(node.Value)
	}

	fmt.Printf("the depth is %d \n", tree.MaxDepth())

	fmt.Println("travel in order")
	tree.TravelInOrder(visit)
	fmt.Println("travel pre order")
	tree.TravelPreOrder(visit)
	fmt.Println("travel post order")
	tree.TravelPostOrder(visit)
}

func TestAddWithMinimalDepth(t *testing.T) {
	nums := []int{}
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	tree := NewBinarySearchTreeWithMinimumDepth(nums)
	fmt.Printf("the depth is %d \n", tree.MaxDepth())
}
