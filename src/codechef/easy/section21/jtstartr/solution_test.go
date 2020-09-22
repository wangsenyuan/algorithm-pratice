package main

import "testing"

func runSample(t *testing.T, n int, V []int, E [][]int, expect int) {
	res := solve(n, V, E)

	if res != expect {
		t.Errorf("Sample %d %v, %v, expect %d, but got %d", n, V, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	V := []int{2, 1, 5, 3}
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
	}
	expect := 2
	runSample(t, n, V, E, expect)
}
