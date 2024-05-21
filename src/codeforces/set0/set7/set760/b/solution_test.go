package main

import "testing"

func runSample(t *testing.T, n int, m, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 10, 3, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 6, 1, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 3, 3, 1)
}