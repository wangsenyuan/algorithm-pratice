package main

import "testing"

func runSample(t *testing.T, n int, arr []int, expect int) {
	res := solve(n, arr)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, []int{1, 2, 3, 1, 2, 3, 8}, 3)
}
