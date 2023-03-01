package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 1},
	}
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{2, 1},
		{2, 3},
		{4, 3},
		{4, 5},
	}
	expect := 0
	runSample(t, n, E, expect)
}
