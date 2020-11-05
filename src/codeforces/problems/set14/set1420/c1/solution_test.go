package main

import "testing"

func runSample(t *testing.T, n int, P []int, expect int64) {
	res := solve(n, P)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := []int{1, 3, 2}
	var expect int64 = 3
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	P := []int{1, 2, 5, 4, 3, 6, 7}
	var expect int64 = 9
	runSample(t, n, P, expect)
}
