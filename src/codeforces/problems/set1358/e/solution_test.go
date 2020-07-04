package main

import "testing"

func runSample(t *testing.T, n int, A []int, X int, expect int) {
	res := solve(n, A, X)

	if res != expect {
		t.Errorf("Sample %d %v %d, expect %d, but got %d", n, A, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{2, -1}
	X := 2
	expect := 2
	runSample(t, n, A, X, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{-2, -2, 6}
	X := -1
	expect := 4
	runSample(t, n, A, X, expect)
}
