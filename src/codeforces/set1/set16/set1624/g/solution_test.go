package main

import "testing"

func runSample(t *testing.T, n int, m int, E [][]int, expect int) {
	res := solve(n, m, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	E := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{1, 3, 2},
	}
	expect := 2
	runSample(t, n, m, E, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 7
	E := [][]int{
		{4, 2, 7},
		{2, 5, 8},
		{3, 4, 2},
		{3, 2, 1},
		{2, 4, 2},
		{4, 1, 2},
		{1, 2, 2},
	}
	expect := 10
	runSample(t, n, m, E, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 4
	E := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{1, 3, 3},
		{3, 1, 4},
	}
	expect := 3
	runSample(t, n, m, E, expect)
}
