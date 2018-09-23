package p909

import "testing"

func runSample(t *testing.T, board [][]int, expect int) {
	res := snakesAndLadders(board)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}
	expect := 4
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := [][]int{
		{-1, 2, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, 52, 8, 66, -1, -1, 78},
		{-1, -1, 31, -1, -1, 67, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, 13, 71, -1},
		{-1, -1, -1, -1, 22, 40, -1, 31, -1},
		{-1, -1, 78, 9, 40, 72, -1, 30, -1},
		{-1, -1, 14, -1, 71, -1, -1, -1, 18},
		{-1, 74, 65, 79, -1, 73, 59, 81, 70},
		{-1, -1, -1, 67, -1, 74, -1, 3, -1},
	}
	expect := 2
	runSample(t, board, expect)
}

func TestSample3(t *testing.T) {
	board := [][]int{
		{-1, -1, -1, 46, 47, -1, -1, -1}, {51, -1, -1, 63, -1, 31, 21, -1}, {-1, -1, 26, -1, -1, 38, -1, -1}, {-1, -1, 11, -1, 14, 23, 56, 57}, {11, -1, -1, -1, 49, 36, -1, 48}, {-1, -1, -1, 33, 56, -1, 57, 21}, {-1, -1, -1, -1, -1, -1, 2, -1}, {-1, -1, -1, 8, 3, -1, 6, 56},
	}
	expect := 4
	runSample(t, board, expect)
}
