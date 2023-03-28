package main

import "testing"

func runSample(t *testing.T, n int, m int, r int, expect int) {
	res := solve(n, m, r)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 3, 6)
}
