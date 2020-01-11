package main

import "testing"

func runSample(t *testing.T, n, k, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 3, 1)
}
