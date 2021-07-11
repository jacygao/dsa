package array

import "testing"

func TestSearchInsertPositionBinarySearch(t *testing.T) {
	testcases := []struct {
		in     []int
		target int
		out    int
	}{
		{
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
		{
			[]int{1, 3, 5, 6},
			0,
			0,
		},
		{
			[]int{1},
			0,
			0,
		},
		{
			[]int{-2, -1, 1, 3, 5, 6},
			0,
			2,
		},
		{
			[]int{-2, -1, 1, 3, 5, 6},
			-3,
			0,
		},
		{
			[]int{-3, -1, 1, 3, 5, 6},
			-2,
			1,
		},
	}

	for _, test := range testcases {
		out := SearchInsertPositionBinarySearch(test.in, test.target)
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SearchInsertPositionBinarySearch([]int{-3, -1, 1, 3, 5, 6}, -2)
	}
}

func BenchmarkSortSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SearchInsertPositionSort([]int{-3, -1, 1, 3, 5, 6}, -2)
	}
}
