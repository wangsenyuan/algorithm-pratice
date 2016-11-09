package main

import "testing"

func TestCountRangeSum(t *testing.T) {
	nums := []int{-2, 5, -1}
	upper := 2
	lower := -2
	expected := 3
	got := countRangeSum(nums, lower, upper)
	if got != expected {
		t.Errorf("expect %d, but get %d\n", expected, got)
	}
}
