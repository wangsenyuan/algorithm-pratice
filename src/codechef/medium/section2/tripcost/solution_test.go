package main

import "testing"

func runSample(t *testing.T, n int, D int, d []int, C []int, a, b int) {
	x, y := solve(n, D, d, C)

	if x != a || y != b {
		t.Errorf("Sample expect %d %d, but got %d %d", a, b, x, y)
	}
}

func TestSample1(t *testing.T) {
	n, D := 4, 5
	d := []int{2, 1, 3, 1}
	c := []int{4, 2, 3, 1}

	a, b := 2, 2

	runSample(t, n, D, d, c, a, b)
}

func TestSample2(t *testing.T) {
	n, D := 5, 6
	d := []int{2, 4, 2, 2, 4}
	c := []int{1, 4, 2, 3, 5}

	a, b := 3, 2

	runSample(t, n, D, d, c, a, b)
}
