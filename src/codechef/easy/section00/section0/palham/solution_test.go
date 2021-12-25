package main

import "testing"

func runSample(t *testing.T, n int, k int, s string, expect int) {
	res := solve(n, k, []byte(s))

	if res != expect {
		t.Errorf("Sample %d %d %s, expect %d, but got %d", n, k, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	s := "abc"
	expect := 76
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 1
	s := "abc"
	expect := 2
	runSample(t, n, k, s, expect)
}
