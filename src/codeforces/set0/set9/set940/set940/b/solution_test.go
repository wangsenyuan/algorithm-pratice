package main

import "testing"

func runSample(t *testing.T, n int, k int, a int, b int, expect int) {
	res := solve(n, k, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 9, 2, 3, 1, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 5, 2, 20, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 19, 3, 4, 2, 12)
}
