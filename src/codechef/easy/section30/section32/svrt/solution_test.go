package main

import "testing"

func runSample(t *testing.T, n int, x int, a, b int) {
	c, d := solve(n, x)

	if c != a || b != d {
		t.Errorf("Sample expect %d, %d, but got %d %d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 3, 2, 3)
}


func TestSample2(t *testing.T) {
	runSample(t, 10, 6, 2, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 4, 3, 2)
}
