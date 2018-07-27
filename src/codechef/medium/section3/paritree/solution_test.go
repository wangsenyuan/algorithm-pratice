package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, m int, conditions [][]int, expect int) {
	res := solve(n, edges, m, conditions)
	if res != expect {
		t.Errorf("Sample %d %v %d %v, expect %d, but got %d", n, edges, m, conditions, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	edges := [][]int{
		{1, 2}, {1, 3},
	}
	conditions := [][]int{
		{1, 2, 0},
		{1, 3, 0},
	}
	expect := 1
	runSample(t, n, edges, m, conditions, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 0
	edges := [][]int{
		{1, 2}, {2, 3},
	}

	expect := 4
	runSample(t, n, edges, m, nil, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 1
	edges := [][]int{
		{1, 2}, {2, 3},
	}
	conditions := [][]int{
		{1, 2, 1},
	}
	expect := 2
	runSample(t, n, edges, m, conditions, expect)
}
