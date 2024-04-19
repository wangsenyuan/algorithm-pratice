package main

import "testing"

func runSample(t *testing.T, x int, y int, k int, expect int) {
	res := solve(x, y, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 3, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3, 7, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 1, 9, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 3, 3, 3)
}
