package main

import "testing"

func runSample(t *testing.T, n int) {

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	ask := func(b int, i int, j int, times int) int {
		grid[i-1][j-1] = b
		return 2
	}

	a := 2

	solve(n, a, ask)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] < 1 || grid[i][j] > 3 {
				t.Fatalf("Sample result got wrong grid %v", grid)
			}
			if i > 0 && grid[i-1][j] == grid[i][j] {
				t.Fatalf("Sample result got wrong grid %v", grid)
			}
			if j > 0 && grid[i][j-1] == grid[i][j] {
				t.Fatalf("Sample result got wrong grid %v", grid)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2)
}


func TestSample2(t *testing.T) {
	runSample(t, 100)
}


func TestSample3(t *testing.T) {
	runSample(t, 11)
}
