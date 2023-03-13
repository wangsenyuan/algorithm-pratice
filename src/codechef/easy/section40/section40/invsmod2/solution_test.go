package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(int64(n), int64(k))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 10, 0
	expect := 1
	runSample(t, n, k, expect)
}
