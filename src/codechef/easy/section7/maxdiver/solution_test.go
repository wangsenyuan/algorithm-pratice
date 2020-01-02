package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int64) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 10
	A := []int{1, 2, 3}
	runSample(t, n, k, A, 3)
}

func TestSample2(t *testing.T) {
	n, k := 4, 2
	A := []int{1, 1, 2, 2}
	runSample(t, n, k, A, 6)
}

func TestSample3(t *testing.T) {
	n, k := 6, 2
	A := []int{2, 3, 3, 2, 4, 4}
	runSample(t, n, k, A, 14)
}
