package main

import "testing"

func runSample(t *testing.T, n int, k int, edges [][]int, expect int) {
	res := solve(n, k, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 5},
		{3, 6},
		{4, 7},
	}
	expect := 7
	runSample(t, n, k, edges, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 1
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}
	expect := 2
	runSample(t, n, k, edges, expect)
}

func TestSample3(t *testing.T) {
	n, k := 8, 5
	edges := [][]int{
		{7, 5},
		{1, 7},
		{6, 1},
		{3, 7},
		{8, 3},
		{2, 1},
		{4, 5},
	}
	expect := 9
	runSample(t, n, k, edges, expect)
}

func TestSample4(t *testing.T) {
	n, k := 20, 7
	edges := [][]int{
		{9, 7},
		{3, 7},
		{15, 9},
		{1, 3},
		{11, 9},
		{18, 7},
		{17, 18},
		{20, 1},
		{4, 11},
		{2, 11},
		{12, 18},
		{8, 18},
		{13, 2},
		{19, 2},
		{10, 9},
		{6, 13},
		{5, 8},
		{14, 1},
		{16, 13},
	}
	expect := 38
	runSample(t, n, k, edges, expect)
}
