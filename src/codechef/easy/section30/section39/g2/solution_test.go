package main

import "testing"

func runSample(t *testing.T, board [][]int, expect int) {
	res := solve(board)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	board := [][]int{
		{2, 1, 1},
		{2, 1, 1},
	}
	expect := 3
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := [][]int{
		{6, 2, 5, 2},
		{4, 5, 3, 8},
		{9, 7, 8, 9},
		{9, 9, 9, 5},
	}
	expect := 5
	runSample(t, board, expect)
}
