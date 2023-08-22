package main

import "testing"

func runSample(t *testing.T, n int64, m int64, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 7, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 5, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 100000, 10000000000, 10000000000)
}
