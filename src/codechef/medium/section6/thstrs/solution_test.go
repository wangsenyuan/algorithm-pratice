package main

import "testing"

func runSample(t *testing.T, a, b, c string, expect int) {
	res := solve(a, b, c)
	if res != expect {
		t.Errorf("Sample %s %s %s, expect %d, but got %d", a, b, c, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abc", "cde", "bccd", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "abc", "cde", "cec", 7)
}
