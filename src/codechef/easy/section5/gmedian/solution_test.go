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
	A := []int{2, 3, 2}
	runSample(t, n, A, 5)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 2, 2, 3, 2}
	runSample(t, n, A, 24)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{1, 2, 2, 2}
	runSample(t, n, A, 12)
}
