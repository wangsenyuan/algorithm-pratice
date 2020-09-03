package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %t, but got %t", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 18, true)
}
