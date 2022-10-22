package main

import "testing"

func runSample(t *testing.T, n int, A []int, edges [][]int, expect int) {
	res := solve(n, A, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	A := []int{1}
	expect := 0
	runSample(t, n, A, nil, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{2, 1, 3}
	edges := [][]int{{1, 3}}
	expect := 2
	runSample(t, n, A, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 3},
	}
	expect := 3
	runSample(t, n, A, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []int{1, 1, 1, 1, 1}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 1},
	}
	expect := 2
	runSample(t, n, A, edges, expect)
}
