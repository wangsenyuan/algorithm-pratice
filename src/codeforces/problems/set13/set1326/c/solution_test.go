package main

import "testing"

func runSample(t *testing.T, n, k int, P []int, expect int64) {
	_, res := solve(n, k, P)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	P := []int{2, 1, 3}
	var expect int64 = 2
	runSample(t, n, k, P, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 5
	P := []int{2, 1, 5, 3, 4}
	var expect int64 = 1
	runSample(t, n, k, P, expect)
}

func TestSample3(t *testing.T) {
	n, k := 7, 3
	P := []int{2, 7, 3, 1, 5, 4, 6}
	var expect int64 = 6
	runSample(t, n, k, P, expect)
}
