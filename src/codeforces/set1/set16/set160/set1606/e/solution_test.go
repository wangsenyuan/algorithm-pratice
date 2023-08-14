package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 5, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 4, 1024)
}

func TestSample4(t *testing.T) {
	runSample(t, 13, 37, 976890680)
}
