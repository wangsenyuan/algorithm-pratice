package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcd", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "aaaa", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "ababcdccdc", 1)
}
