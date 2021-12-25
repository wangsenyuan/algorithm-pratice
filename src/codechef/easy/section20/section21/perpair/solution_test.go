package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 3, 6)
}
