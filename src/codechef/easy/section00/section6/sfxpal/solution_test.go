package main

import "testing"

func runSample(t *testing.T, n, s, m int, expect int) {
	res := solve(n, s, m)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", n, s, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 13, 12)
}
