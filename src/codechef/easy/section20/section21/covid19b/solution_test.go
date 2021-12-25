package main

import "testing"

func runSample(t *testing.T, n int, V []int, a, b int) {
	c, d := solve(n, V)

	if c != a || d != b {
		t.Errorf("Sample expect %d %d, but got %d %d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	V := []int{1, 2, 3}
	a, b := 1, 1
	runSample(t, n, V, a, b)
}

func TestSample2(t *testing.T) {
	n := 3
	V := []int{3, 2, 1}
	a, b := 3, 3
	runSample(t, n, V, a, b)
}

func TestSample3(t *testing.T) {
	n := 3
	V := []int{0, 0, 0}
	a, b := 1, 1
	runSample(t, n, V, a, b)
}

func TestSample4(t *testing.T) {
	n := 3
	V := []int{1, 3, 2}
	a, b := 1, 2
	runSample(t, n, V, a, b)
}
