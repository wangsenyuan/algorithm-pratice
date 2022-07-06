package main

import "testing"

func runSample(t *testing.T, m int, n int, k int, expect int) {
	res := solve(k, m, n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n, k := 2, 2, 5
	expect := 14
	runSample(t, m, n, k, expect)
}

func TestSample2(t *testing.T) {
	m, n, k := 2, 3, 7
	expect := 5
	runSample(t, m, n, k, expect)
}
