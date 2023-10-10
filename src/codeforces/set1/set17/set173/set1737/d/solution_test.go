package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	E := [][]int{
		{1, 2, 3},
		{6, 4, 5},
		{3, 5, 6},
		{6, 1, 3},
		{7, 4, 4},
		{3, 8, 4},
		{2, 3, 3},
		{7, 8, 5},
		{4, 5, 2},
	}
	expect := 9
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	E := [][]int{
		{4, 6, 92},
		{7, 1, 65},
		{6, 5, 43},
		{6, 7, 96},
		{4, 3, 74},
		{4, 8, 54},
		{7, 4, 99},
		{2, 5, 22},
	}
	expect := 154
	runSample(t, n, E, expect)
}
