package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, 3},
		{1, 3, 5},
	}
	expect := 49
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 3},
		{1, 3, 5},
		{2, 4, 7},
		{2, 5, 2},
	}
	expect := 174
	runSample(t, n, E, expect)
}
