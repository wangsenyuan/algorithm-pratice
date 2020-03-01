package main

import "testing"

func runSample(t *testing.T, N, M, L int, B []int, edges [][]int, expect int) {
	res := solve(N, M, L, B, edges)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", N, M, L, B, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, L := 3, 3, 6
	B := []int{1, 2, 3, 1, 2, 3}
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{3, 1, 1},
	}
	expect := 6

	runSample(t, N, M, L, B, edges, expect)
}

func TestSample2(t *testing.T) {
	N, M, L := 4, 4, 9
	B := []int{1, 2, 3, 4, 1, 2, 3, 4, 1}
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{3, 4, 1},
		{4, 1, 1},
	}
	expect := 5

	runSample(t, N, M, L, B, edges, expect)
}

func TestSample3(t *testing.T) {
	N, M, L := 3, 3, 2
	B := []int{1, 2}
	edges := [][]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 1},
	}
	expect := -1

	runSample(t, N, M, L, B, edges, expect)
}
