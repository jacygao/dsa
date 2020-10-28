package main

import (
	"flag"
	"log"

	"github.com/jacygao/dsa/linkedlist"
)

// Problem statement
//
// Write code to partition a linked list around a value x, such that
// all nodes less than x come before all nodes greater than or equial
// to x. If x is contained within th elist, the values of x only need
// to be after the elements less than x (see below). The partition element
// x can appear anywhere in the "right partition"; it does not need to
// appear between the left and right partitions.
//

var input int

var data = []int{
	7, 10, 4, 22, 8, 10, 7, 15, 1, 10, 8, 40,
}

var llist *linkedlist.LinkedList

// Solve runs logic to solve the problem.
// This algorithm iterates the list only once.
// If it finds any node with value greater than the input value, it shifts
// the nodes to the end of the linked list.
// If the node value is less than the input no change needs to be made to the
// node.
func Solve(v int) {
	node := llist.First
	walker := node
	var firstEndNode *linkedlist.Node
	counter := 0

	for walker != firstEndNode {
		node = walker
		walker = node.Next()
		if node.Value().(int) >= v {
			llist.ShiftToEnd(node)
			if counter == 0 {
				firstEndNode = node
				counter++
			}
		}
	}

	printResult()
}

func printResult() {
	var resultList []int

	node := llist.First
	for node != nil {
		resultList = append(resultList, node.Value().(int))
		node = node.Next()
	}

	log.Printf("result: %v", resultList)
}

// init gets executed first to set up the prerequisites.
func init() {
	// Assume the linklist to made of a list of integers
	log.Printf("list: %v", data)

	llist = linkedlist.New()
	for _, d := range data {
		llist.Append(d)
	}
}

func main() {
	flag.IntVar(&input, "v", 0, "the partition value")
	flag.Parse()

	if input == 0 {
		panic("missing input")
	}

	Solve(input)
}
