package graph

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 2},
		{0, 3},
		{3, 1},
		{4, 1},
		{4, 2},
		{5, 0},
		{5, 2},
	}

	expected := []int{5, 4, 0, 3, 1, 2}
	res := TopologicalSort(n, edges)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("results do not match! expcted %+v but got %+v", expected, res)
	}
}
