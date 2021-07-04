package strint

import (
	"testing"
)

func TestReverseInt(t *testing.T) {
	testcases := []struct {
		in int
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
			10,
			0,
		},
		{
			-123,
			0,
		},
		{
			2356712351,
			1532176532,
		},
	}

	for _, test := range testcases {
		out := ReverseInt(test.in) 
		if out != test.out {
			t.Fatalf("expected %d, but got %d", test.out, out)
		}
	}
}