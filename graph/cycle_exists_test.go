package graph

import "testing"

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
