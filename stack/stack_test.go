package stack

import (
	"reflect"
	"testing"
)

func TestPushPopPeekIsEmpty(t *testing.T) {
	s := New()

	val1 := "testdata1"
	val2 := "testdata2"

	s.Push(val1)
	s.Push(val2)

	val3 := s.Pop()
	if !reflect.DeepEqual(val3, val2) {
		t.Fatalf("expected %v but got %v", val2, val3)
	}

	if !reflect.DeepEqual(s.Peek(), val1) {
		t.Fatalf("expected %v but got %v", val1, s.Peek())
	}

	_ = s.Pop()

	if !s.IsEmpty() {
		t.Fatal("not expecting an non-empty stack")
	}
}
