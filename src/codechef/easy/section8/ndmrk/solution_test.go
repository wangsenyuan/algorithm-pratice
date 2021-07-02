package main

import "testing"

func runSample(t *testing.T, n, k int, P []int, expect int) {
	res := solve(n, k, P)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 5
	P := []int{1, 1}
	expect := 4
	runSample(t, n, k, P, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 6
	P := []int{1, 2, 2}
	expect := 3
	runSample(t, n, k, P, expect)
}
