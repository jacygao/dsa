package array

import (
	"fmt"
	"sort"
	"testing"
)

type CombinationSumInput struct {
	candidates []int
	target     int
}

func TestCombinationSumSol1(t *testing.T) {
	testcases := []struct {
		in  CombinationSumInput
		out [][]int
	}{
		{
			CombinationSumInput{
				[]int{2, 3, 6, 7},
				7,
			},
			[][]int{
				[]int{2, 2, 3},
				[]int{7},
			},
		},
		{
			CombinationSumInput{
				[]int{2, 3, 5},
				8,
			},
			[][]int{
				[]int{2, 2, 2, 2},
				[]int{2, 3, 3},
				[]int{3, 5},
			},
		},
		{
			CombinationSumInput{
				[]int{1},
				1,
			},
			[][]int{
				[]int{1},
			},
		},
		{
			CombinationSumInput{
				[]int{1},
				2,
			},
			[][]int{
				[]int{1, 1},
			},
		},
		{
			CombinationSumInput{
				[]int{2, 7, 6, 3, 5, 1},
				9,
			},
			[][]int{
				[]int{1, 1, 1, 1, 1, 1, 1, 1, 1},
				[]int{1, 1, 1, 1, 1, 1, 1, 2},
				[]int{1, 1, 1, 1, 1, 1, 3},
				[]int{1, 1, 1, 1, 1, 2, 2},
				[]int{1, 1, 1, 1, 2, 3},
				[]int{1, 1, 1, 1, 5},
				[]int{1, 1, 1, 2, 2, 2},
				[]int{1, 1, 1, 3, 3},
				[]int{1, 1, 1, 6},
				[]int{1, 1, 2, 2, 3},
				[]int{1, 1, 2, 5},
				[]int{1, 1, 7},
				[]int{1, 2, 2, 2, 2},
				[]int{1, 2, 3, 3},
				[]int{1, 2, 6},
				[]int{1, 3, 5},
				[]int{2, 2, 2, 3},
				[]int{2, 2, 5},
				[]int{2, 7},
				[]int{3, 3, 3},
				[]int{3, 6},
			},
		},
	}

	for _, test := range testcases {
		out := CombinationSumSol1(test.in.candidates, test.in.target)
		if !isEqual(out, test.out) {
			t.Fatalf("expected %v, but got %v", test.out, out)
		}
	}
}

func TestIsEqual(t *testing.T) {
	if !isEqual(
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		},
		[][]int{
			[]int{3, 2, 1},
			[]int{5, 4, 6},
		},
	) {
		t.Fatal("expected 2 slices of slice to be equal!")
	}
}

func isEqual(s1, s2 [][]int) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[string]struct{}, len(s1))
	for _, val1 := range s1 {
		sort.Ints(val1)
		m[fmt.Sprintf("%v", val1)] = struct{}{}
	}

	for _, val2 := range s2 {
		sort.Ints(val2)
		if _, ok := m[fmt.Sprintf("%v", val2)]; !ok {
			return false
		}
	}
	return true
}
