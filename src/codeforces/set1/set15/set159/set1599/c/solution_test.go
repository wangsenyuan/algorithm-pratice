package main

import "testing"

func runSample(t *testing.T, n int, p float64, expect int) {
	res := solve(n, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 1.0, 6)
}
