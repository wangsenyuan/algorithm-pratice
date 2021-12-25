package main

import "testing"

func runSample(t *testing.T, n int, m int, E [][]int, expect int) {
	res := solve(n, m, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 6, 5
	E := [][]int{
		{1, 6},
		{2, 4},
		{2, 5},
		{3, 5},
		{4, 5},
	}
	expect := 1
	runSample(t, n, m, E, expect)
}
