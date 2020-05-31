package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	A := []int{1}
	runSample(t, n, A, 0)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{2, 1}
	runSample(t, n, A, 0)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{4, 3, 2, 1}
	runSample(t, n, A, 1)
}

func TestSample4(t *testing.T) {
	n := 6
	A := []int{5, 4, 3, 2, 6, 1}
	runSample(t, n, A, 1)
}
