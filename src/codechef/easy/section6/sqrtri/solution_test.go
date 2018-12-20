package main

import "testing"

func runSample(t *testing.T, a, b, x, y int, expect bool) {
	res := solve(a, b, x, y)

	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %t, but got %t", a, b, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 1, 1, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1, 1, 2, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 1, 0, 0, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 2, 1, 1, false)
}
