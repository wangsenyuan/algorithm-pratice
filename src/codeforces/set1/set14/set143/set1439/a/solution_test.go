package main

import "testing"

func runSample(t *testing.T, grid []string) {
	res := solve(grid)

	n := len(grid)
	m := len(grid[0])
	maze := make([][]byte, n)

	for i := 0; i < n; i++ {
		maze[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				maze[i][j] = 0
			} else {
				maze[i][j] = 1
			}
		}
	}

	for _, cur := range res {
		for i := 0; i < 6; i += 2 {
			r, c := cur[i]-1, cur[i+1]-1
			maze[r][c] ^= 1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maze[i][j] != 0 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"10",
		"11",
	}
	runSample(t, grid)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"011",
		"101",
		"110",
	}
	runSample(t, grid)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"1111",
		"0110",
		"0110",
		"1111",
	}
	runSample(t, grid)
}
