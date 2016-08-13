package main

import "testing"

func TestPlay(t *testing.T) {
	tests := []struct {
		nums []int
		m    int
		want bool
	}{
		{[]int{1, 1, 1}, 3, true},
		{[]int{1, 2, 4, 8, 16}, 11, true},
		{[]int{1, 2, 4, 8, 16}, 23, true},
		{[]int{1, 5, 5, 10, 10}, 13, false},
		{[]int{17, 6, 4, 998, 254, 137, 259, 153, 154, 3, 28, 19, 123, 542, 857, 23, 687, 35, 99, 999}, 132, true},
	}

	for _, test := range tests {
		got := play(test.m, test.nums)
		if got != test.want {
			t.Errorf("play(%v) got %v is wrong, it should be %v\n", test.nums, got, test.want)
		}
	}
}
