package main

import "testing"

func runSample(t *testing.T, n int, d int, expect int) {
	res := solve(n, d)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 21, 5, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 8, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 0, 11)
}

func TestSample4(t *testing.T) {
	runSample(t, 5925, 9, 75)
}

func TestSample5(t *testing.T) {
	runSample(t, 434356, 3, 5644)
}
