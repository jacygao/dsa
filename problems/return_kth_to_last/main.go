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
// This algorithm walks the linked list to get the length of the linked list.
// Then we substract the input value to get the distence from the end of the list.
func Solve(k int) {
	length := 0

	node := llist.First
	for node != nil {
		node = node.Next()
		length++
	}

	output := length - input

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
