package main

import "testing"

func runSample(t *testing.T, s string, k, x int, expect int) {
	res := solve(s, k, x)

	if res != expect {
		t.Errorf("Sample %s %d %d, expect %d, but got %d", s, k, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcdefagahai", 0, 1, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "abcdefagahai", 1, 1, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, "abcdefagahai", 2, 1, 8)
}
