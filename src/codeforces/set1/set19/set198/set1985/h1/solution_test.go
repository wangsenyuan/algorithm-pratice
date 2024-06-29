package main

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := solve(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		".",
	}
	expect := 1
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"..",
		"#.",
		"#.",
		".#",
	}
	expect := 6
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		".#.#.",
		"..#..",
		".#.#.",
	}
	expect := 9
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{
		"#...#",
		"....#",
		"#...#",
		".....",
		"...##",
	}
	expect := 11
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := []string{
		"..#....#",
		".####.#.",
		"###.#..#",
		".##.#.##",
		".#.##.##",
		"#..##.#.",
	}
	expect := 30
	runSample(t, grid, expect)
}
