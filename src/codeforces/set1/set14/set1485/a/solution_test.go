package main

import "testing"

func runSample(t *testing.T, a, b, c int) {
	d := solve(a, b)
	if c != d {
		t.Errorf("Sample %d %d, expect %d, but got %d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 9, 2, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 1337, 1, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, 991026972, 997, 3)
}

func TestSample5(t *testing.T) {
	runSample(t, 1, 1, 2)
}
