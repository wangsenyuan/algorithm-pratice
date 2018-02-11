package p782

import "testing"

func runSample(t *testing.T, board [][]int, expect int) {
	res := movesToChessboard(board)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := [][]int{
		{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1},
	}
	runSample(t, board, 2)
}

func TestSample2(t *testing.T) {
	board := [][]int{
		{0, 1}, {1, 0},
	}
	runSample(t, board, 0)
}

func TestSample3(t *testing.T) {
	board := [][]int{
		{0, 1}, {0, 1},
	}
	runSample(t, board, -1)
}
