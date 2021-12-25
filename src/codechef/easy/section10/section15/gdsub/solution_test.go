package main

import "testing"

func runSmaple(t *testing.T, n, k int, arr []int, expect int64) {
	res := solve(n, k, arr)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 3
	arr := []int{2, 2, 3, 3, 5}
	expect := int64(18)
	runSmaple(t, n, k, arr, expect)
}
