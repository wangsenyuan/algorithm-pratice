package main

import "testing"

func runSample(t *testing.T, n, k int, E [][]int, expect int) {
	res := solve(n, k, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 3
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 8
	runSample(t, n, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 2
	E := [][]int{
		{1, 6},
		{2, 4},
		{2, 6},
		{3, 6},
		{5, 6},
	}
	expect := 25
	runSample(t, n, k, E, expect)
}
