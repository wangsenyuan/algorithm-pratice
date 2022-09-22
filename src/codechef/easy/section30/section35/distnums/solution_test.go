package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 5, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, 30)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 1, 180)
}
