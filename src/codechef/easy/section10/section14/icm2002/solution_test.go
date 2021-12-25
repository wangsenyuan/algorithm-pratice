package main

import "testing"

func runSample(t *testing.T, n, k int, B []int, expect bool) {
	res := solve(n, k, B)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, k, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 3
	B := []int{5, 4, 3, 2, 1}
	runSample(t, n, k, B, true)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	B := []int{4, 1, 1}
	runSample(t, n, k, B, false)
}
