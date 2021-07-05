package strint

import (
	"testing"
)

func TestReverseInt(t *testing.T) {
	testcases := []struct {
		in  int
		out int
	}{
		{
			123,
			321,
		},
		{
			0,
			0,
		},
		{
			-7,
			-7,
		},
		{
			210,
			12,
		},
		{
			-123,
			-321,
		},
		{
			1356712359,
			0,
		},
	}

	for _, test := range testcases {
		out := ReverseInt(test.in)
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}
