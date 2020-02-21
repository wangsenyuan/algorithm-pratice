package main

import "testing"

func runSample(t *testing.T, n, m, k int, A []int, E [][]int, expect int) {
	res := solve(n, m, k, A, E)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", n, m, k, A, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 5, 5, 3
	A := []int{1, 3, 5}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{2, 4},
	}
	expect := 3
	runSample(t, n, m, k, A, E, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 5, 4, 2
	A := []int{2, 4}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 3
	runSample(t, n, m, k, A, E, expect)
}
