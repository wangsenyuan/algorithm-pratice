package main

import "testing"

func TestSample(t *testing.T) {
	n := 5
	l, r := int64(1), int64(4)
	nums := []int64{1, 2, 3, 4, 5}
	res := solve(n, l, r, nums)
	if res != 3 {
		t.Errorf("the sample should give anser 3, but got %d", res)
	}
}
