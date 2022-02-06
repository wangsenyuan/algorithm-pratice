package main

import "testing"

func runSample(t *testing.T, n, m, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 2, 2, 2
	expect := 10
	runSample(t, n, m, k, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 2, 3, 2
	expect := 22
	runSample(t, n, m, k, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 41, 42, 2
	expect := 903408624
	runSample(t, n, m, k, expect)
}
