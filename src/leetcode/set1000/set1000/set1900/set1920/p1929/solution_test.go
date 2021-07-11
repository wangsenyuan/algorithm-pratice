package p1929

import (
	"testing"
)

func runSample(t *testing.T, maze []string, entr []int, expect int) {
	grid := make([][]byte, len(maze))
	for i := 0; i < len(maze); i++ {
		grid[i] = []byte(maze[i])
	}
	res := nearestExit(grid, entr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	maze := []string{
		"++.+", "...+", "+++.",
	}
	entr := []int{1, 2}
	expect := 1
	runSample(t, maze, entr, expect)
}

func TestSample2(t *testing.T) {
	maze := []string{
		"+++", "...", "+++",
	}
	entr := []int{1, 0}
	expect := 2
	runSample(t, maze, entr, expect)
}
