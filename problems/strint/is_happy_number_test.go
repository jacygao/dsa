package strint

import "testing"

func TestIsHappyNumber(t *testing.T) {
	testcases := []struct {
		in  int
		out bool
	}{
		{
			19,
			true,
		},
		{
			1,
			true,
		},
		{
			1000000,
			true,
		},
		{
			2,
			false,
		},
		{
			1234,
			false,
		},
	}

	for _, test := range testcases {
		out := IsHappyNumber(test.in)
		if out != test.out {
			t.Fatalf("expected %t, but got %t", test.out, out)
		}
	}
}
