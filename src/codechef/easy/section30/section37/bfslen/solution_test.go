package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := int(solve(n, E))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 6
	runSample(t, n, E, expect)
}
