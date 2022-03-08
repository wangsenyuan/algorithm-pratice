package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := make([][]int, 0)
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{3, 2},
		{1, 4},
		{5, 3},
		{4, 2},
	}
	expect := 0
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{3, 2},
		{1, 3},
	}
	expect := 3
	runSample(t, n, E, expect)
}
