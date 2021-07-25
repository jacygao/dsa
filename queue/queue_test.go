package queue

import (
	"reflect"
	"testing"
)

func TestAddPeekRemoveIsEmpty(t *testing.T) {
	q := New()

	val1 := "testdata1"
	val2 := "testdata2"

	q.Add(val1)
	q.Add(val2)

	peekVal := q.Peek()

	// testing FIFO
	if !reflect.DeepEqual(peekVal, val1) {
		t.Fatalf("expect peek item to be %v but got %v", val1, peekVal)
	}

	if q.IsEmpty() {
		t.Fatal("expect queue to contain 2 bot got empty")
	}

	q.Remove()
	q.Remove()

	if !q.IsEmpty() {
		t.Fatal("expect queue to be empty")
	}
}
