package p1906

import "testing"

func runSample(t *testing.T, grid1, grid2 [][]int, expect int) {
	res := countSubIslands(grid1, grid2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid1 := [][]int{
		{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1},
	}
	grid2 := [][]int{
		{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0},
	}
	expect := 3
	runSample(t, grid1, grid2, expect)
}

func TestSample2(t *testing.T) {
	grid1 := [][]int{
		{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1},
	}
	grid2 := [][]int{
		{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1},
	}
	expect := 2
	runSample(t, grid1, grid2, expect)
}
