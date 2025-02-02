package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 2)
}

func TestSample3(t *testing.T) {
	// 2 + 2
	runSample(t, 2, 2, 6)
}
