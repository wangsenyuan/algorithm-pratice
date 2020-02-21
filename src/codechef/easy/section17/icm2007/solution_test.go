package main

import "testing"

func runSample(t *testing.T, n, m int, E [][]int, expect bool) {
	res := solve(n, m, E)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, m, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 4
	E := [][]int{
		{1, 2, 2},
		{2, 3, 1},
		{4, 1, -1},
		{3, 4, -2},
	}
	expect := true

	runSample(t, n, m, E, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 4
	E := [][]int{
		{1, 2, 2},
		{2, 3, 2},
		{1, 4, 2},
		{4, 3, 1},
	}
	expect := false

	runSample(t, n, m, E, expect)
}
