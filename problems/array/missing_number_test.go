package array

import "testing"

func TestMissingNumberOptimal(t *testing.T) {
	testcases := []struct {
		in  []int
		out int
	}{
		{
			[]int{3, 0, 1},
			2,
		},
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			8,
		},
		{
			[]int{0},
			1,
		},
	}

	for _, test := range testcases {
		out := MissingNumberOptimal(test.in)
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}

func TestMissingNumberSort(t *testing.T) {
	testcases := []struct {
		in  []int
		out int
	}{
		{
			[]int{3, 0, 1},
			2,
		},
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			8,
		},
		{
			[]int{0},
			1,
		},
	}

	for _, test := range testcases {
		out := MissingNumberSort(test.in)
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}

var testdata = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

func BenchmarkMissingNumberOptimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MissingNumberOptimal(testData)
	}
}

func BenchmarkMissingNumberSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MissingNumberSort(testData)
	}
}
