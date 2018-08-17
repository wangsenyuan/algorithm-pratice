package main

import "testing"

func runSample(t *testing.T, R, C int, puzzle []string, expect int) {
	res := solve(R, C, puzzle)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", R, C, puzzle, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C := 5, 4
	puzzle := []string{
		"....",
		"#..#",
		"#xx#",
		"#oo#",
		"#..#",
	}
	expect := 2
	runSample(t, R, C, puzzle, expect)
}

func TestSample2(t *testing.T) {
	R, C := 7, 7
	puzzle := []string{
		".######",
		".x....#",
		".x....#",
		"..#oo.#",
		"..#...#",
		".######",
		".######",
	}
	expect := 8
	runSample(t, R, C, puzzle, expect)
}

func TestSample3(t *testing.T) {
	R, C := 7, 7
	puzzle := []string{
		"#######",
		"#x....#",
		"#xx.o.#",
		"#..oo.#",
		"#.#...#",
		"#######",
		"#######",
	}
	expect := 12
	runSample(t, R, C, puzzle, expect)
}

func TestSample4(t *testing.T) {
	R, C := 4, 10
	puzzle := []string{
		"##########",
		"#.x...o..#",
		"#.x...o..#",
		"##########",
	}
	expect := 8
	runSample(t, R, C, puzzle, expect)
}

func TestSample5(t *testing.T) {
	R, C := 7, 7
	puzzle := []string{
		"#######",
		"#x....#",
		"#x..o.#",
		"#x#oo.#",
		"#.#...#",
		"#######",
		"#######",
	}
	expect := -1
	runSample(t, R, C, puzzle, expect)
}

func TestSample6(t *testing.T) {
	R, C := 3, 4
	puzzle := []string{
		".#x.",
		".oww",
		"....",
	}
	expect := 2
	runSample(t, R, C, puzzle, expect)
}

func TestSample7(t *testing.T) {
	R, C := 3, 4
	puzzle := []string{
		".#x.",
		".o..",
		"....",
	}
	expect := 2
	runSample(t, R, C, puzzle, expect)
}

func TestSample8(t *testing.T) {
	R, C := 3, 4
	puzzle := []string{
		".#x.",
		"o...",
		"....",
	}
	expect := -1
	runSample(t, R, C, puzzle, expect)
}

func TestSample9(t *testing.T) {
	R, C := 12, 10
	puzzle := []string{
		"####..###.",
		"..##..#...",
		"##...#...#",
		".###.##..#",
		"##....##..",
		".##..o.#.#",
		"..#..oooo#",
		"####......",
		".#....#.##",
		"#..xxx....",
		"##...xx..#",
		"##.###.#..",
	}
	expect := 56
	runSample(t, R, C, puzzle, expect)
}
