package main

import "testing"

func runSample(t *testing.T, n, p, k int, expect int) {
	res := solve(n, p, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 5, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 6, 5, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 4, 5, 9)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 8, 5, 8)
}