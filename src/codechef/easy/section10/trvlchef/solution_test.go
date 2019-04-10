package main

import "testing"

func runSample(t *testing.T, n int, D int, C []int, expect bool) {
	res := solve(n, D, C)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, D, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	D := 3
	C := []int{3, 2, 1, 4, 5}
	expect := true
	runSample(t, n, D, C, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	D := 4
	C := []int{10, 1, 3, 2, 9}
	expect := false
	runSample(t, n, D, C, expect)
}
