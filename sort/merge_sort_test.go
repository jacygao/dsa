package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{5, 7, 6, 1, 2, 10, 8, 3, 4, 9}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := MergeSort(input)
	if !reflect.DeepEqual(input, expected) {
		t.Fatalf("results do not match! expected %v but got %v", expected, res)
	}
}
