package main

import "testing"

func runSample(t *testing.T, n, x, y int, expect int) {
	res := solve(n, x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 3, 1, 0)
}
