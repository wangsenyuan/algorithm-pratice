package main

import "testing"

func runSample(t *testing.T, n int, l int64, r int64, z int64, expect int) {
	res := solve(n, l, r, z)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 5, 1, 13)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 1, 3, 2, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1, 100000, 15629, 49152)
}

func TestSample4(t *testing.T) {
	runSample(t, 100, 56, 89, 66, 981727503)
}
