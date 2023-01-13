package main

import "testing"

func runSample(t *testing.T, h int, w int, x int, y int) {
	a, b := solve(h, w)

	if a != x || b != y {
		t.Errorf("Sample expect %d, %d, but got %d %d", x, y, a, b)
	}
}

func TestSample1(t *testing.T) {
	h, w := 2, 1
	x, y := 1, 1
	runSample(t, h, w, x, y)
}

func TestSample2(t *testing.T) {
	h, w := 2, 2
	x, y := 2, 2
	runSample(t, h, w, x, y)
}

func TestSample3(t *testing.T) {
	h, w := 5, 5
	x, y := 5, 4
	runSample(t, h, w, x, y)
}

func TestSample4(t *testing.T) {
	h, w := 9, 10
	x, y := 8, 10
	runSample(t, h, w, x, y)
}

func TestSample5(t *testing.T) {
	h, w := 4774, 4806
	x, y := 4096, 4806
	runSample(t, h, w, x, y)
}
