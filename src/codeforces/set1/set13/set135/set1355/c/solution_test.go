package main

import "testing"

func runSample(t *testing.T, a int, b int, c int, d int, expect int) {
	res := solve(a, b, c, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 1 3 3
	// 2 2 3
	// 2 3 3
	// 2 3 4
	runSample(t, 1, 2, 3, 4, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, 2, 5, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 500000, 500000, 500000, 500000, 1)
}
