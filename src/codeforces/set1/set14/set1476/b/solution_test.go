package main

import "testing"

func runSample(t *testing.T, n int, k int, P []int, expect int64) {
	res := solve(n, k, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 1
	P := []int{20100, 1, 202, 202}
	var expect int64 = 99
	runSample(t, n, k, P, expect)
}
