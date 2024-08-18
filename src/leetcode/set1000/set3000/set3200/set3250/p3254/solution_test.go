package p3254

import "testing"

func runSample(t *testing.T, board [][]int, expect int) {
	res := maximumValueSum(board)

	if expect != int(res) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := [][]int{
		{-3, 1, 1, 1},
		{-3, 1, -3, 1},
		{-3, 2, 1, 1},
	}
	expect := 4
	runSample(t, board, expect)
}

func TestSample2(t *testing.T) {
	board := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := 15
	runSample(t, board, expect)
}

func TestSample3(t *testing.T) {
	board := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	expect := 3
	runSample(t, board, expect)
}

func TestSample4(t *testing.T) {
	board := [][]int{
		{5, 29, 12},
		{23, -37, 99},
		{12, 31, 57},
	}
	expect := 140
	runSample(t, board, expect)
}
