package main

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{2, 1, 2}, 3},
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{5, 4, 1, 2}, 8},
	}

	for _, test := range tests {
		got := largestRectangleArea(test.heights)
		if got != test.want {
			t.Errorf("%d != %d for %v\n", got, test.want, test.heights)
		}
	}
}
