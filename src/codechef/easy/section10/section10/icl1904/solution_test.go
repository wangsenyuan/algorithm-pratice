package main

import "testing"

func runSample(t *testing.T, N uint64, S string, expect int) {
	res := solve(N, S)

	if res != expect {
		t.Errorf("%d %s, expect %d, but got %d", N, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, "a", 1326)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, "ab", 76)
}
