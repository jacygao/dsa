package array

import (
	"reflect"
	"testing"
)

type Input struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func TestMergeSortedArray(t *testing.T) {
	testcases := []struct {
		in  Input
		out []int
	}{
		{
			Input{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			Input{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			[]int{1},
		},
		{
			Input{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			[]int{1},
		},
	}

	for _, test := range testcases {
		MergeSortedArray(test.in.nums1, test.in.m, test.in.nums2, test.in.n)
		if !reflect.DeepEqual(test.in.nums1, test.out) {
			t.Fatalf("expected %v, but got %v", test.out, test.in.nums1)
		}
	}
}

var testdata2 = Input{
	nums1: []int{1, 2, 3, 0, 0, 0},
	m:     3,
	nums2: []int{2, 5, 6},
	n:     3,
}

// A single loop with time complexity of O(n)
func BenchmarkMergeWithForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortedArray(testdata2.nums1, testdata2.m, testdata2.nums2, testdata2.n)
	}
}

// Sorting has time complexity of n log(n)
func BenchmarkMergeWithSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeAndSort(testdata2.nums1, testdata2.m, testdata2.nums2, testdata2.n)
	}
}
