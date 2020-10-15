package main

import "testing"

func runSample(t *testing.T, n int, A []int, a, b int) {
	c, d := solve(n, A)

	if a != c || b != d {
		t.Errorf("Sample %d %v, expect %d %d, but got %d %d", n, A, a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{2, 9, 8, 2, 7}
	a, b := 2, 3
	runSample(t, n, A, a, b)
}

func TestSample2(t *testing.T) {
	n := 1
	A := []int{2}
	a, b := 1, 0
	runSample(t, n, A, a, b)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{1, 1, 1}
	a, b := 2, 1
	runSample(t, n, A, a, b)
}
