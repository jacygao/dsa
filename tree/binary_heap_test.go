package tree

import (
	"reflect"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	heap := NewMinHeap(10)
	for i := 1; i < 11; i++ {
		heap.Insert(i, 11-i)
	}

	res := []int{}
	// priority is the opposite of value in test data
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for heap.size != 0 {
		res = append(res, heap.ExtractMin().val)
	}

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("results do not match! expected %+v but got %+v", expected, res)
	}
}

func TestBinaryHeapSmoke(t *testing.T) {
	heap := NewMinHeap(5)
	heap.Insert(3, 5)
	heap.Insert(2, 10)
	heap.Insert(7, 5)
	heap.Insert(1, 100)
	heap.Insert(14, 9000)
	heap.Insert(0, 500)

	res := []int{}
	// priority is the opposite of value in test data
	expected := []int{100, 10, 5, 5, 9000}

	for heap.size != 0 {
		res = append(res, heap.ExtractMin().val)
	}

	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("results do not match! expected %+v but got %+v", expected, res)
	}
}
