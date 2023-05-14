package p2677

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maxMoves(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{187, 167, 209, 251, 152, 236, 263, 128, 135},
		{267, 249, 251, 285, 73, 204, 70, 207, 74},
		{189, 159, 235, 66, 84, 89, 153, 111, 189},
		{120, 81, 210, 7, 2, 231, 92, 128, 218},
		{193, 131, 244, 293, 284, 175, 226, 205, 245},
	}
	expect := 3
	runSample(t, grid, expect)
}
