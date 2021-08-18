package array

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	testcases := []struct {
		in  []int
		out []int
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0},
			[]int{0},
		},
	}

	for _, test := range testcases {
		MoveZeroesSol1(test.in)
		if !reflect.DeepEqual(test.out, test.in) {
			t.Fatalf("expected %d, but got %d", test.out, test.in)
		}
	}
}

var moveZeroesTestdata = []int{0, 1, 0, 3, 0, 12, 0, 45, 2, 0, 0, 8}

func BenchmarkMoveZeroesSol1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MoveZeroesSol1(moveZeroesTestdata)
	}
}

func BenchmarkMoveZeroesSol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MoveZeroesSol2(moveZeroesTestdata)
	}
}

func BenchmarkMoveZeroesSol3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MoveZeroesSol3(moveZeroesTestdata)
	}
}
