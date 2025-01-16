package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 49, 1808, 359087121)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 9, 30)
}
