package main

import "testing"

func TestSecondLargest(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 9, 8, 7}
	got := secondLargest(nums)
	if got != 8 {
		t.Errorf("second largest in %v should be %d, but is %d\n", nums, 8, got)
	}
}
