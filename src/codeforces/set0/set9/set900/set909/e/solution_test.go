package main

import "testing"

func runSample(t *testing.T, n int, deps [][]int, E []int, expect int) {
	res := solve(n, deps, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	E := []int{0, 1, 0, 1}
	n := 4
	deps := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}
	expect := 2
	runSample(t, n, deps, E, expect)
}

func TestSample2(t *testing.T) {
	E := []int{1, 1, 1, 0}
	n := 4
	deps := [][]int{
		{0, 1},
		{0, 2},
		{3, 0},
	}
	expect := 1
	runSample(t, n, deps, E, expect)
}
