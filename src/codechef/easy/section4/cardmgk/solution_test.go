package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect bool) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, []int{1, 5, 2, 4, 3}, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, []int{3, 4, 5, 1, 2}, true)
}
