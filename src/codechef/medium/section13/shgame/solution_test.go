package main

import "testing"

func runSample(t *testing.T, n, m int, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 8, 40)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 7, 42)
}
