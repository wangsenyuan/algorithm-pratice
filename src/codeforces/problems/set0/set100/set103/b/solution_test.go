package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{6, 3},
		{6, 4},
		{5, 1},
		{2, 5},
		{1, 4},
		{5, 4},
	}
	expect := true

	runSample(t, n, E, expect)
}
