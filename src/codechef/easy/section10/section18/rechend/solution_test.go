package main

import "testing"

func runSample(t *testing.T, n int, C [][]int, expect bool) {
	res := solve(n, C)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	C := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	expect := true
	runSample(t, n, C, expect)
}
