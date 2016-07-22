package main

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{[]int{0, 0}, 1},
	}

	for _, test := range tests {
		got := wiggleMaxLength(test.nums)
		if got != test.want {
			t.Errorf("wiggleMaxLength(%v) => %v not correct with %d\n", test.nums, got, test.want)
		}
	}
}
