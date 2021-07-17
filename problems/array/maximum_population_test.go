package array

import "testing"

func TestMaximumPopulation(t *testing.T) {
	testcases := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{[]int{1993, 1999}, []int{2000, 2010}},
			1993,
		},
		{
			[][]int{[]int{1950, 1961}, []int{1960, 1971}, []int{1970, 1981}},
			1960,
		},
		{
			[][]int{
				[]int{2033, 2034},
				[]int{2039, 2047},
				[]int{1998, 2042},
				[]int{2047, 2048},
				[]int{2025, 2029},
				[]int{2005, 2044},
				[]int{1990, 1992},
				[]int{1952, 1956},
				[]int{1984, 2014},
			},
			2005,
		},
	}

	for _, test := range testcases {
		out := MaximumPopulation(test.in)
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}
