package p773

import "testing"

func runSample(t *testing.T, board [][]int, expect int) {
	res := slidingPuzzle(board)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := [][]int{{1, 2, 3}, {4, 0, 5}}
	runSample(t, board, 1)
}

func TestSample2(t *testing.T) {
	board := [][]int{{1, 2, 3}, {5, 4, 0}}
	runSample(t, board, -1)
}

func TestSample3(t *testing.T) {
	board := [][]int{{4, 1, 2}, {5, 0, 3}}
	runSample(t, board, 5)
}

func TestSample4(t *testing.T) {
	board := [][]int{{3, 2, 4}, {1, 5, 0}}
	runSample(t, board, 14)
}
