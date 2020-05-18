package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{2, 3, 6}
	runSample(t, n, A, 1)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{4, 4, 8}
	runSample(t, n, A, 2)
}
