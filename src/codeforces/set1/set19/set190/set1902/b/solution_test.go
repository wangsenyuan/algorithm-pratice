package main

import "testing"

func runSample(t *testing.T, n int, P int, l int, x int, expect int) {
	res := solve(n, P, l, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 5, 5, 2, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 14, 3000000000, 1000000000, 500000000, 12)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 20, 1, 10, 99)
}
