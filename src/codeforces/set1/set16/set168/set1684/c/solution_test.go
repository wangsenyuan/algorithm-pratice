package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, board [][]int, expect bool) {
	res := solve(board)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	x, y := res[0], res[1]
	x--
	y--
	for i := 0; i < len(board); i++ {
		board[i][x], board[i][y] = board[i][y], board[i][x]
		if !sort.IntsAreSorted(board[i]) {
			t.Fatalf("Sample result %v, not making %d-th row sorted", res, i)
		}
	}
}
func TestSample1(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{1, 1, 1},
	}
	expect := true

	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := [][]int{
		{4, 1},
		{2, 3},
	}
	expect := false

	runSample(t, board, expect)
}

func TestSample3(t *testing.T) {
	board := [][]int{
		{2, 1},
		{1, 1},
	}
	expect := true

	runSample(t, board, expect)
}

func TestSample4(t *testing.T) {
	board := [][]int{
		{6, 2, 1},
		{5, 4, 3},
	}
	expect := true

	runSample(t, board, expect)
}
