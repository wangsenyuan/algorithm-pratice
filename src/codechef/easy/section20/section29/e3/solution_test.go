package main

import "testing"

func runSample(t *testing.T, n int, P []int, expect int) {
	res := solve(n, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	P := []int{0, 0, 1, 0, 0, 1, 1, 2, 1, 1}
	expect := 2
	runSample(t, n, P, expect)
}
