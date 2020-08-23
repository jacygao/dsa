package linkedlist

import (
	"reflect"
	"testing"
)

func TestAppendOneElement(t *testing.T) {
	ll := New()
	var data interface{}
	data = "mock"
	ll.Append(data)

	if !reflect.DeepEqual(ll.First.Value(), data) {
		t.Fatalf("incorrect result! expected %+v but got %+v", ll.First.Value(), data)
	}

	if !reflect.DeepEqual(ll.Last.Value(), data) {
		t.Fatalf("incorrect result! expected %+v but got %+v", ll.Last.Value(), data)
	}
}

func TestAppendMultipleElement(t *testing.T) {
	ll := New()
	strList := []interface{}{
		"a",
		"b",
		"c",
		"d",
	}
	for _, v := range strList {
		ll.Append(v)
	}

	cur := ll.First
	for _, v := range strList {
		if !reflect.DeepEqual(cur.Value(), v) {
			t.Fatalf("incorrect result! expected %+v but got %+v", cur.Value(), v)
		}
		cur = cur.Next()
	}
}

func TestDeleteNode(t *testing.T) {
	ll := New()
	strList := []interface{}{
		"a",
		"b",
		"c",
		"d",
		"e",
	}
	for _, v := range strList {
		ll.Append(v)
	}

	secondNode := ll.First.Next()
	ll.Delete(secondNode)
	if !reflect.DeepEqual(ll.First.Next().Value(), strList[2]) {
		t.Fatalf("incorrect result! expected %+v but got %+v", strList[2], ll.First.Next().Value())
	}

	firstNode := ll.First
	ll.Delete(firstNode)
	if !reflect.DeepEqual(ll.First.Value(), strList[2]) {
		t.Fatalf("incorrect result! expected %+v but got %+v", strList[2], ll.First.Value())
	}

	lastNode := ll.Last
	ll.Delete(lastNode)
	if !reflect.DeepEqual(ll.Last.Value(), strList[3]) {
		t.Fatalf("incorrect result! expected %+v but got %+v", strList[3], ll.Last.Value())
	}
}
