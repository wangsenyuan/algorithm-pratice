package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, test := range tests {
		got := climbStairs(test.n)
		if got != test.want {
			t.Errorf("climbStairs(%d) => %d is wrong with %d", test.n, test.want, got)
		}
	}
}
