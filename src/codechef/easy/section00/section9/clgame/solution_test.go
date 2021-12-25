package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, S []int, expect bool) {
	res := solve(n, k, A, S)
	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %t, but got %t", n, k, A, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	A := []int{5, 7, 1}
	S := []int{1, 2}
	runSample(t, n, k, A, S, true)
}
