package main

import "testing"

func runSample(t *testing.T, n int, A []int, k int, expect int) {
	res := solve(n, A, k)
	if res != expect {
		t.Errorf("Sample %d %v %d, expect %d, but got %d", n, A, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	A := []int{1, 1, 3, 3, 5}
	expect := 4
	runSample(t, n, A, k, expect)
}
