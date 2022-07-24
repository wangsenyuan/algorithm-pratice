package main

import "testing"

func runSample(t *testing.T, n int, k int, s int, expect int) {
	res := solve(n, k, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, s := 16, 2, 10
	expect := 2
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k, s := 50, 48, 7
	expect := -1
	runSample(t, n, k, s, expect)
}
