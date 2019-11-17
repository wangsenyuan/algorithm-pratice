package p1263

import "testing"

func runSample(t *testing.T, grid [][]string, expect int) {
	board := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		board[i] = make([]byte, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			board[i][j] = grid[i][j][0]
		}
	}
	res := minPushBox(board)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]string{
		{"#", "#", "#", "#", "#", "#"},
		{"#", "T", "#", "#", "#", "#"},
		{"#", ".", ".", "B", ".", "#"},
		{"#", ".", "#", "#", ".", "#"},
		{"#", ".", ".", ".", "S", "#"},
		{"#", "#", "#", "#", "#", "#"},
	}
	runSample(t, grid, 3)
}

func TestSample2(t *testing.T) {
	grid := [][]string{
		{"#", "#", "#", "#", "#", "#"},
		{"#", "T", "#", "#", "#", "#"},
		{"#", ".", ".", "B", ".", "#"},
		{"#", "#", "#", "#", ".", "#"},
		{"#", ".", ".", ".", "S", "#"},
		{"#", "#", "#", "#", "#", "#"},
	}
	runSample(t, grid, -1)
}

func TestSample3(t *testing.T) {
	grid := [][]string{
		{"#", "#", "#", "#", "#", "#"},
		{"#", "T", ".", ".", "#", "#"},
		{"#", ".", "#", "B", ".", "#"},
		{"#", ".", ".", ".", ".", "#"},
		{"#", ".", ".", ".", "S", "#"},
		{"#", "#", "#", "#", "#", "#"},
	}
	runSample(t, grid, 5)
}

func TestSample4(t *testing.T) {
	grid := [][]string{
		{"#", "#", "#", "#", "#", "#", "#"},
		{"#", "S", "#", ".", "B", "T", "#"},
		{"#", "#", "#", "#", "#", "#", "#"},
	}
	runSample(t, grid, -1)
}
