package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int, expect int) {
	res := int(solve(n, E, A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{2, 1},
		{1, 3},
		{4, 3},
		{3, 5},
	}
	A := []int{2, 5}
	expect := 5
	runSample(t, n, E, A, expect)
}
