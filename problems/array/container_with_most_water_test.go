package array

import "testing"

func TestMaxArea2Ptr(t *testing.T) {
	testcases := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{4, 3, 2, 1, 4},
			16,
		},
		{
			[]int{1, 2, 1},
			2,
		},
	}

	for _, test := range testcases {
		out := MaxArea2Ptr(test.in)
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}

var testData = []int{1, 8, 6, 2, 5, 4, 8, 3, 7, 5, 3, 4, 6, 7, 8, 6, 5, 4, 3, 20, 13, 5, 2, 3, 4, 5, 5}

func BenchmarkMaxArea2Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxArea2Ptr(testData)
	}
}
