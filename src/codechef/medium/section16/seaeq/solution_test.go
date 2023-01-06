package main

import "testing"

func runSample(t *testing.T, n int, E int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 0, 9)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 1, 1)
}
