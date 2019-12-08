package main

import "testing"

func runSample(t *testing.T, n int, k int, A string, expect int) {
	res := solve(n, k, A)
	if res != expect {
		t.Errorf("Sample %d %d %s, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 2, "0101110000101", 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 3, "100001", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 4, "100001", 6)
}
