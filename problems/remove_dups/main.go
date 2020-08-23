package main

import (
	"log"

	"github.com/jacygao/linkedlist"
)

// Problem statement
//
// Write code to remove duplicates from an unsorted linked list.
// How would you solve this problem if a temprory buffer is not allowed?
//

var llist *linkedlist.LinkedList

// Solve runs logic to solve the problem.
// This algorithm will take event node from head, comparen with each following node.
// If any duplicate value is found, the node is deleted.
// The process stops when it reaches the end of the linked list.
func Solve() {
	target := llist.First
	for target.Next() != nil {
		walker := target.Next()
		for walker != nil {
			if match(target, walker) {
				newTarget := target.Next()
				llist.Delete(target)
				target = newTarget
				break
			}
			walker = walker.Next()
		}

		// walker has reached the end of the linked list and no match has been found.
		// Current node has no duplicate value and let's move targe to the next node.
		if walker == nil {
			target = target.Next()
		}
	}

	printResult()
}

func match(target, walker *linkedlist.Node) bool {
	if target.Value() == walker.Value() {
		return true
	}
	return false
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
	data := []int{
		7, 10, 4, 22, 8, 10, 7, 15, 1, 10, 8, 40,
	}

	log.Printf("list: %v", data)

	llist = linkedlist.New()
	for _, d := range data {
		llist.Append(d)
	}
}

func main() {
	Solve()
}
