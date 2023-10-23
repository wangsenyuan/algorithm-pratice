package main

import "testing"

func runSample(t *testing.T, n int, p int, expect int) {
	res := solve(n, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 100000007, 16)
}

func TestSample2(t *testing.T) {
	runSample(t, 1145, 141919831, 105242108)
}
