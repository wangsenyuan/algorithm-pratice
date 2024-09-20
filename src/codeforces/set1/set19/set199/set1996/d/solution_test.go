package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 4, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 5, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 1000, 7)
}

func TestSample4(t *testing.T) {
	runSample(t, 900000, 400000, 1768016938)
}
