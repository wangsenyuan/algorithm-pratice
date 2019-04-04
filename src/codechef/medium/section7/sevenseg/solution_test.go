package main

import "testing"

func runSample(t *testing.T, n int, X []int, Y []int, a int, b int) {
	c, d := solve(n, X, Y)
	if a > b {
		// invalid
		if c < d {
			t.Errorf("Sample %d %v %v, expect invalid answer, but got %d %d", n, X, Y, c, d)
		}
	} else {
		// valid
		if a != c || b != d {
			t.Errorf("Sample %d %v %v, expect %d %d, but got %d %d", n, X, Y, a, b, c, d)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 2
	X := []int{2, 5}
	Y := []int{4, 4}
	a, b := 1, 2
	runSample(t, n, X, Y, a, b)
}

func TestSample2(t *testing.T) {
	n := 2
	X := []int{8, 0}
	Y := []int{6, 7}
	a, b := 3, 2
	runSample(t, n, X, Y, a, b)
}
