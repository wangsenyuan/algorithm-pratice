package main

import "testing"

func runSample(t *testing.T, maze []string, expect int) {
	res := solve(maze)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	maze := []string{
		".#.",
		"...",
		".#.",
	}
	expect := 2
	runSample(t, maze, expect)
}

func TestSample2(t *testing.T) {
	maze := []string{
		"##.",
		"...",
		".#.",
		".#.",
	}
	expect := 2
	runSample(t, maze, expect)
}

func TestSample3(t *testing.T) {
	maze := []string{
		"..##",
		"....",
		"..#.",
	}
	expect := 2
	runSample(t, maze, expect)
}

func TestSample4(t *testing.T) {
	maze := []string{
		"##..",
		"..#.",
		"...#",
		"...#",
	}
	expect := -1
	runSample(t, maze, expect)
}

func TestSample5(t *testing.T) {
	maze := []string{
		".##..",
		".##..",
		".#.#.",
		"..#..",
		"..#..",
	}
	expect := 2
	runSample(t, maze, expect)
}

func TestSample6(t *testing.T) {
	maze := []string{
		"###",
		"###",
		"###",
	}
	expect := 2
	runSample(t, maze, expect)
}
