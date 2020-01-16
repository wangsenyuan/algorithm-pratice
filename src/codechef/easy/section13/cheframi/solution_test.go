package main

import "testing"

func runSample(t *testing.T, n int, x int, A []int, expect int64) {
	res := solve(n, x, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, x, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 3, 3
	A := []int{5, 2, 4}
	runSample(t, n, x, A, 8)
}
