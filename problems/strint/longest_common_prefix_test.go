package strint

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testcases := []struct {
		in  []string
		out string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			[]string{""},
			"",
		},
		{
			[]string{"dog", "racecar", "car"},
			"",
		},
		{
			[]string{"abc"},
			"abc",
		},
		{
			[]string{"ab", "a"},
			"a",
		},
	}

	for _, test := range testcases {
		out := LongestCommonPrefix(test.in)
		if out != test.out {
			t.Fatalf("expected %s, but got %s", test.out, out)
		}
	}
}
