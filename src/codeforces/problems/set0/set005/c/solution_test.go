package main

import "testing"

func runSample(t *testing.T, s string, a, b int) {
	c, d := solve(s)

	if c != a || b != d {
		t.Errorf("Sample %s, expect %d %d, but got %d %d", s, a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	s := ")((())))(()())"
	a, b := 6, 2
	runSample(t, s, a, b)
}

