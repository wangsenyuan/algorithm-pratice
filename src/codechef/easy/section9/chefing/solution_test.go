package main

import "testing"

func runSample(t *testing.T, n int, S []string, expect int) {
	res := solve(n, S)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	S := []string{"abcaa", "bcbd", "bgc"}
	expect := 2
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	S := []string{"quick", "brown", "fox"}
	expect := 0
	runSample(t, n, S, expect)
}
