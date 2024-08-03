package main

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	cnt, res := solve(grid)

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}
	var tmp int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != res[i][j] {
				tmp++
			}
			if i > 0 && res[i][j] == res[i-1][j] || j > 0 && res[i][j] == res[i][j-1] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
	if tmp != expect {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"aaaa",
		"bbbb",
		"cccc",
	}
	expect := 6
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"aba",
		"aba",
		"zzz",
	}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"aba",
		"aba",
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{
		"a",
	}
	expect := 0
	runSample(t, grid, expect)
}
