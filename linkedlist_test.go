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
		t.Fatalf("incorrect result! expected %+v but got %+v", data, ll.First.Value())
	}

	if ll.First.Next() != nil {
		t.Fatalf("incorrect result! expected %+v but got %+v", nil, ll.First.Next())
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

	lastNode := ll.Last()
	ll.Delete(lastNode)
	if !reflect.DeepEqual(ll.Last().Value(), strList[3]) {
		t.Fatalf("incorrect result! expected %+v but got %+v", strList[3], ll.Last().Value())
	}
}

func TestLastNode(t *testing.T) {
	ll := New()
	if ll.Last() != nil {
		t.Fatal("expected empty linked list has nil last node")
	}

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

	if ll.Last().Value().(string) != strList[4] {
		t.Fatalf("expected %+v but got %+v", ll.Last().Value(), strList[4])
	}
}

func TestShiftToEnd(t *testing.T) {
	testcase := []struct {
		desc     string
		nodeth   int
		expFirst string
		expLast  string
	}{
		{
			"nil node changes nothing",
			0,
			"a",
			"e",
		},
		{
			"shift to end the first node",
			1,
			"b",
			"a",
		},
		{
			"shift to end the last node",
			2,
			"a",
			"e",
		},
		{
			"shift to end a middle node",
			3,
			"a",
			"b",
		},
	}

	for _, tt := range testcase {
		t.Run(tt.desc, func(t *testing.T) {
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

			var node *Node
			switch tt.nodeth {
			case 0:
				node = nil
				break
			case 1:
				node = ll.First
				break
			case 2:
				node = ll.Last()
			case 3:
				node = ll.First.Next()
			default:
				t.Fatal("unsupported nodeth value")
			}

			ll.ShiftToEnd(node)

			firstVal := ll.First.Value().(string)
			if firstVal != tt.expFirst {
				t.Fatalf("expected %+v but got %+v", tt.expFirst, firstVal)
			}

			lastVal := ll.Last().Value()
			if lastVal != tt.expLast {
				t.Fatalf("expected %+v but got %+v", tt.expLast, lastVal)
			}
		})
	}
}
