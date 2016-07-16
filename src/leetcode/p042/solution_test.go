package main

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{2, 0, 2}, 2},
	}

	for _, test := range tests {
		got := trap(test.height)
		if got != test.want {
			t.Errorf("%d is not the anwer %d for %v\n", got, test.want, test.height)
		}
	}
}
