package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 4, 2, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 1, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 4, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 10, 2, 30)
}
