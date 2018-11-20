package main

import "testing"

func runSample(t *testing.T, n, m int, edges [][]int, expect int) {
	res := solve(n, m, edges)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	expect := 3
	runSample(t, n, m, edges, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 2
	edges := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := 1
	runSample(t, n, m, edges, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 0
	edges := [][]int{}
	expect := 0
	runSample(t, n, m, edges, expect)
}
