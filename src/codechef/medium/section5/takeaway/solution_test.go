package main

import "testing"

func runSample(t *testing.T, n, k int, stones []int, expect bool) {
	res := solve(n, k, stones)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, k, stones, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 1, 1
	stones := []int{3}
	expect := true
	runSample(t, n, k, stones, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 2
	stones := []int{3, 3}
	expect := false
	runSample(t, n, k, stones, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 10
	stones := []int{5, 10, 15, 20, 25}
	expect := true
	runSample(t, n, k, stones, expect)
}
