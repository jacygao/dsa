package array

import "testing"

func TestFindShortestSubArray(t *testing.T) {
	testcases := []struct {
		in  []int
		out int
	}{
		{
			[]int{0},
			1,
		},
		{
			[]int{1, 1, 1},
			3,
		},
		{
			[]int{6, 5, 5},
			2,
		},
		{
			[]int{1, 2, 2, 3, 1},
			2,
		},
		{
			[]int{1, 2, 2, 3, 1, 4, 2},
			6,
		},
	}

	for _, test := range testcases {
		out := FindShortestSubArray(test.in)
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}

func BenchmarkFindShortestSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FindShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})
	}
}
