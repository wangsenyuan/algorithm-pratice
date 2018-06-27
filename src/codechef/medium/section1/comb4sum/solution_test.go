package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 2, 3, 3}
	runSample(t, n, A, 10)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	runSample(t, n, A, 76)
}

func TestSample3(t *testing.T) {
	n := 6
	A := []int{9, 18, 28, 23, 12, 9}
	runSample(t, n, A, 1176)
}
