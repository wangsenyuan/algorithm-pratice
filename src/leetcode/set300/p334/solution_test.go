package main

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		nums   []int
		expect bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 3, 2, 4, 2}, true},
	}

	for _, test := range tests {
		got := increasingTriplet(test.nums)
		if got != test.expect {
			t.Errorf("should increasingTriplet(%v) == %v, but get %v\n", test.nums, test.expect, got)
		}
	}
}
