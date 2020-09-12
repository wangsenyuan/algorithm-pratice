package main

import "testing"

func runSample(t *testing.T, x, y, n int, expect int64) {
	res := solve(x, y, n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 5, 12345, 12339)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 0, 4, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 5, 15, 15)
}

func TestSample4(t *testing.T) {
	runSample(t, 17, 8, 54321, 54306)
}

func TestSample5(t *testing.T) {
	runSample(t, 499999993, 9, 1000000000, 999999995)
}

func TestSample6(t *testing.T) {
	runSample(t, 2, 0, 999999999, 999999998)
}
