package main

import "testing"

func runSample(t *testing.T, n int, X []int, expect int64) {
	res := solve(n, X)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	X := []int{-1, 1}
	runSample(t, n, X, 0)
}

func TestSample2(t *testing.T) {
	n := 1
	X := []int{0, 0}
	runSample(t, n, X, 1)
}

func TestSample3(t *testing.T) {
	n := 2
	X := []int{4, 3, 1, 4}
	runSample(t, n, X, 2)
}

func TestSample4(t *testing.T) {
	n := 3
	X := []int{5, 3, 7, 10, 5, 10}
	runSample(t, n, X, 4)
}
