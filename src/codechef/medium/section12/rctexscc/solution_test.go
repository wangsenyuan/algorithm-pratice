package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, m, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 2, 873463811)
}
