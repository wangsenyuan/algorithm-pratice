package main

import "testing"

func runSample(t *testing.T, n int, A []int, k int, x int, expect int64) {
	res := solve(n, A, k, x)
	if res != expect {
		t.Errorf("Sample %d %v %d %d, expect %d, but got %d", n, A, k, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	k := 2
	x := 4
	runSample(t, n, A, k, x, 23)
}

func TestSample2(t *testing.T) {
	n := 7
	A := []int{10, 15, 20, 13, 2, 1, 44}
	k := 4
	x := 14
	runSample(t, n, A, k, x, 129)
}
