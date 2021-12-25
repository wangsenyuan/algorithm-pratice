package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int) {
	res := solve(n, k, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 6, 2
	A := []int{1, 1, 1, 2, 2, 1}
	expect := 3
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 3
	A := []int{1, 1, 2, 2, 1}
	expect := 5
	runSample(t, n, k, A, expect)
}
