package main

import "testing"

func runSample(t *testing.T, n int, k int, l int, expect int64) {
	res := solve(n, k, l)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, l := 3, 1, 1
	expect := 3
	runSample(t, n, k, l, int64(expect))
}
