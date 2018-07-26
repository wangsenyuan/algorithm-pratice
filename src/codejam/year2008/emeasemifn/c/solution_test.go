package main

import "testing"

func runSample(t *testing.T, n, k int, edges [][]int, expect int) {
	res := solve(n, k, edges)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 10
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 720
	runSample(t, n, k, edges, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 3
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 6
	runSample(t, n, k, edges, expect)
}

func TestSample3(t *testing.T) {
	n, k := 6, 10
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
		{5, 6},
	}
	expect := 40320
	runSample(t, n, k, edges, expect)
}

func TestSample4(t *testing.T) {
	n, k := 7, 10
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
		{5, 6},
		{5, 7},
	}
	expect := 282240
	runSample(t, n, k, edges, expect)
}
