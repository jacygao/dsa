package graph

import (
	"reflect"
	"testing"
)

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
		{1, 2},
		{2, 0},
		{1, 3},
		{3, 4},
		{4, 5},
		{5, 3},
	}
	expected := [][]int{
		{1, 3},
	}
	ans := findBridges(n, edges)
	if !reflect.DeepEqual(ans, expected) {
		t.Fatalf("results do not match! expected %+v but got %+v", expected, ans)
	}
}
