package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect uint64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{4, 3, 5, 1, 2}
	var expect uint64 = 11
	runSample(t, n, A, expect)
}
