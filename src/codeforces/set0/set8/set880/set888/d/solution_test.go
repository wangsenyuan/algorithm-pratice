package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 3, 31)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 4, 76)
}
