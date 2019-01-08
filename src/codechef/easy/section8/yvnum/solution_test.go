package main

import "testing"

func runSample(t *testing.T, N string, expect int) {
	res := solve(N)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "123", 123231312)
}
