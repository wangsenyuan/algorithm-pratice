package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 1, 16)
}
