package main

import "testing"

func runSample(t *testing.T, n, m, a int, expect int64) {
	res := solve(n, m, a)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, m, a, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 6, 4, 4)
}
