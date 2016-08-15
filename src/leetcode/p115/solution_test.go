package main

import "testing"

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		{"ccc", "c", 3},
		{"ddd", "dd", 3},
	}

	for _, test := range tests {
		got := numDistinct(test.s, test.t)
		if got != test.want {
			t.Errorf("there should be %d distinct sub sequence for %v from %v, but get %d\n", test.want, test.t, test.s, got)
		}
	}
}
