package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 5, 13)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000, 1, 1000000000000000000)
}
