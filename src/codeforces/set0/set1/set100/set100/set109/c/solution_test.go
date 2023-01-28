package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := int(solve(n, E))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2, 4},
		{3, 1, 2},
		{1, 4, 7},
	}
	expect := 16
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2, 4},
		{3, 1, 7},
		{1, 4, 7},
	}
	expect := 24
	runSample(t, n, E, expect)
}
