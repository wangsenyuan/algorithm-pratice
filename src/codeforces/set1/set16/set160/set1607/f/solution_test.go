package main

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := solve(grid)

	if res[2] != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	n := len(grid)
	m := len(grid[0])

	r, c := res[0], res[1]
	r--
	c--
	for i := 0; i < expect; i++ {
		if r < 0 || r == n || c < 0 || c == m {
			t.Fatalf("Sample result from %v, cant move %d steps", res[:2], expect)
		}

		if grid[r][c] == 'L' {
			c--
		} else if grid[r][c] == 'R' {
			c++
		} else if grid[r][c] == 'U' {
			r--
		} else {
			r++
		}
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"R",
	}
	expect := 1
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"RRL",
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"DL",
		"RU",
	}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{
		"UD",
		"RU",
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := []string{
		"DL",
		"UL",
		"RU",
	}
	expect := 5
	runSample(t, grid, expect)
}

func TestSample6(t *testing.T) {
	grid := []string{
		"RRRD",
		"RUUD",
		"URUD",
		"ULLR",
	}
	expect := 12
	runSample(t, grid, expect)
}
