package p035

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		index  int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}

	for _, test := range tests {
		nums := test.nums
		target := test.target
		want := test.index
		got := searchInsert(nums, target)
		if want != got {
			t.Errorf("%d != %d for inserting %d into %v\n", want, got, target, nums)
		}
	}
}
