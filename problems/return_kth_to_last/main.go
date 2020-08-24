package main

import (
	"flag"
	"log"

	"github.com/jacygao/linkedlist"
)

var input int

// Problem statement
//
// Implement an algorithm to find the kth element to last element of a
// singly linked list.
//

var data = []int{
	7, 10, 4, 22, 8, 10, 7, 15, 1, 10, 8, 40,
}

var llist *linkedlist.LinkedList

// Solve runs logic to solve the problem.
// This algorithm walks the linked list with 2 pointers.
// The lead pointer is exactly k node ahead of the tail pointer.
// When the lead pointer reaches the end of the list, return the value of
// the tail pointer.
func Solve(k int) {
	counter := 0
	leadNode := llist.First
	tailNode := llist.First

	for leadNode != nil {
		leadNode = leadNode.Next()
		if counter >= k {
			tailNode = tailNode.Next()
		}
		counter++
	}

	output := tailNode.Value()

	log.Printf("result: %d", output)
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
	flag.IntVar(&input, "k", 0, "the kth element")
	flag.Parse()

	if input == 0 {
		panic("missing input")
	}

	if input > len(data) {
		panic("input is out of index")
	}

	Solve(input)
}
