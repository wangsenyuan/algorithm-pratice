package main

import "testing"

func runSample(t *testing.T, n int, W, C []int, expect int64) {
	res := solve(n, W, C)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, W, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	C := []int{0, 1, 2, 0, 2}
	W := []int{5, 6, 7, 8, 2}
	runSample(t, n, W, C, 21)
}
