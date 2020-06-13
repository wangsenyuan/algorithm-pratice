package main

import "testing"

func runSample(t *testing.T, n, k int, X, D []int, expect int) {
	res := solve(n, k, X, D)

	if res != expect {
		t.Errorf("Sample %d %d, %v %v, expect %d, but got %d", n, k, X, D, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 0
	X := []int{50, 40, 30}
	D := []int{2, 2, 3}
	expect := 20
	runSample(t, n, k, X, D, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 1
	X := []int{50, 40, 30}
	D := []int{2, 2, 3}
	expect := 30
	runSample(t, n, k, X, D, expect)
}
