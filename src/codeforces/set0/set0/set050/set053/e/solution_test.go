package main

import "testing"

func runSample(t *testing.T, n int, k int, E [][]int, expect int) {
	res := solve(n, k, E)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := 3
	runSample(t, n, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
		{1, 3},
		{2, 4},
	}
	expect := 12
	runSample(t, n, k, E, expect)
}

func TestSample3(t *testing.T) {
	n, k := 4, 3
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
		{1, 3},
		{2, 4},
	}
	expect := 4
	runSample(t, n, k, E, expect)
}