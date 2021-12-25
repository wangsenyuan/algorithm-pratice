package main

import "testing"

func runSample(t *testing.T, n int, A []int, edges [][]int, expect int) {
	res := solve(n, A, edges)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{0, 1, 1}
	edges := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 1
	runSample(t, n, A, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{0, 1, 2}
	edges := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 1
	runSample(t, n, A, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{2, 2, 2, 2}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
	}
	expect := 0
	runSample(t, n, A, edges, expect)
}
