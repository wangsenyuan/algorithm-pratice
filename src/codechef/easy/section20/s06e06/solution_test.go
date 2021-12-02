package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int64) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2221, 3, 36)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 5, 6)
}