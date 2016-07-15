package main

import "testing"

func TestMaxSubSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-1}, -1},
	}

	for _, test := range tests {
		got := maxSubArray(test.nums)
		if got != test.want {
			t.Errorf("%d != %d max sum array (at least one) in array %v\n", got, test.want, test.nums)
		}
	}
}
