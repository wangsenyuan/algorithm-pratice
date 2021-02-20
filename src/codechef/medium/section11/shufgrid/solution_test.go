package main

import "testing"

func runSample(t *testing.T, n, k, r, c int, expect int) {
	res := solve(n, k, r, c)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, r, c := 1, 2, 1, 1
	expect := 1
	runSample(t, n, k, r, c, expect)
}

func TestSample2(t *testing.T) {
	n, k, r, c := 2, 1, 1, 1
	expect := 249561089
	runSample(t, n, k, r, c, expect)
}
