package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := solve(n, s)

	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, "abab", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, "acbba", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, "aaaab", -1)
}
