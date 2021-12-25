package main

import "testing"

func runSample(t *testing.T, n int, P []int, expect int64) {
	res := solve(n, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := []int{1, 1}
	var expect int64 = 4
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	P := []int{1, 1, 2, 2}
	var expect int64 = 9
	runSample(t, n, P, expect)
}
