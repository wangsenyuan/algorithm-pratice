package main

import "testing"

func runSample(t *testing.T, n int, X []int, P []int, expect int) {
	res := solve(n, X, P)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, X, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	X := []int{1, 2, 3, 4, 5}
	P := []int{1, 2, 3, 4, 5}
	expect := 27
	runSample(t, n, X, P, expect)
}
