package main

import (
	"log"

	"github.com/jacygao/dsa/linkedlist"
)

// Problem statement
//
// You have two numbers represented by a linked list, where each node
// contains a single digit. The digits are stored in reverse order,
// such that the 1's digit is at the head of the list. Write a function
// that adds the two numbers and returns the sum as a linked list.
//

// Solve runs logic to solve the problem.
// This algorithm loops both linked list once while adding new nodes to
// the output linked list which has the time complexity of O(n)
// Each loop will get the sum of two numbers from each linked list.
// If the sum is greater than 10 it carries 1 to the next loop.
// The loop exits when both linked list reaches the last node.
func Solve(l1, l2 *linkedlist.LinkedList) *linkedlist.LinkedList {
	if l1 == nil && l2 == nil {
		return nil
	}

	out := linkedlist.New()

	node1 := l1.First
	node2 := l2.First
	carry := 0

	for node1 != nil || node2 != nil {
		val1 := 0
		val2 := 0

		if node1 != nil {
			val1 = node1.Value().(int)
		}
		if node2 != nil {
			val2 = node2.Value().(int)
		}

		sum := val1 + val2 + carry
		carry = 0

		if sum > 9 {
			sum = sum % 10
			carry = 1
		}
		out.Append(sum)

		if node1 != nil {
			node1 = node1.Next()
		}
		if node2 != nil {
			node2 = node2.Next()
		}
	}

	return out
}

func print(ll *linkedlist.LinkedList, name string) {
	var resultList []int

	node := ll.First
	for node != nil {
		resultList = append(resultList, node.Value().(int))
		node = node.Next()
	}

	log.Printf("Linked List %s: %v", name, resultList)
}

func main() {
	s1 := []int{5, 0, 9, 2, 7}
	s2 := []int{3, 0, 6}
	l1 := linkedlist.New()
	l2 := linkedlist.New()

	for _, s := range s1 {
		l1.Append(s)
	}
	for _, s := range s2 {
		l2.Append(s)
	}

	print(l1, "input 1")
	print(l2, "input 2")

	out := Solve(l1, l2)
	print(out, "result")
}
