package main

import "testing"

func runSample(t *testing.T, n int, x int, y int, expect int) {
	res := solve(n, x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 3, 1, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 4, 3, 44)
}
