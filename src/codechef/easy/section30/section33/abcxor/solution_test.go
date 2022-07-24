package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := int(solve(n, k))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 1
	expect := 6
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 0
	expect := 42
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 9, 100
	expect := 260610
	runSample(t, n, k, expect)
}
