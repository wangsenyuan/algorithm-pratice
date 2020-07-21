package main

import "testing"

func runSample(t *testing.T, R, C int, r, c int, board [][]int, expect int) {
	res := solve(R, C, r, c, board)

	if res != expect {
		t.Errorf("Sample %d %d %d %d %v, expect %d, but got %d", R, C, r, c, board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C, r, c := 3, 3, 2, 2
	board := [][]int{
		{2, 3, 2},
		{1, 7, 1},
		{2, 1, 2},
	}
	expect := INF

	runSample(t, R, C, r, c, board, expect)
}

func TestSample2(t *testing.T) {
	R, C, r, c := 3, 4, 1, 2
	board := [][]int{
		{1, 2, 2, 0},
		{10, 8, 5, 10},
		{10, 2, 9, 10},
	}
	expect := 3

	runSample(t, R, C, r, c, board, expect)
}

func TestSample4(t *testing.T) {
	R, C, r, c := 3, 5, 3, 4
	board := [][]int{
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{1, 1, 1, 1, 0},
	}
	expect := 1
	runSample(t, R, C, r, c, board, expect)
}
