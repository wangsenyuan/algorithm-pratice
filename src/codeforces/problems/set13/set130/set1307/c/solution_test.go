package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aaabb", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "usaco", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "lol", 2)
}
