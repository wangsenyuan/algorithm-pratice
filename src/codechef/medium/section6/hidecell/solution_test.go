package main

import "testing"

func runSample(t *testing.T, n int, cell []int) {
	a, b := cell[0], cell[1]
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	ask := func(grid [][]byte) int {
		tmp := grid[a][b]
		grid[a][b] = '1'
		res := check(grid, dp)
		grid[a][b] = tmp
		return res
	}

	res := solve(n, ask)

	if res[0] != cell[0] || res[1] != cell[1] {
		t.Errorf("Sample expect %v, but got %v", cell, res)
	}
}

func check(grid [][]byte, dp [][]int) int {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}
	if grid[0][0] != '1' || grid[n-1][n-1] != '1' {
		return 0
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 1 {
				if j+1 < n && grid[i][j+1] == '1' {
					dp[i][j+1] = 1
				}
				if i+1 < n && grid[i+1][j] == '1' {
					dp[i+1][j] = 1
				}
			}
		}
	}
	return dp[n-1][n-1]
}

func TestSample1(t *testing.T) {
	n := 4
	cell := []int{2, 1}
	runSample(t, n, cell)
}

func TestSample2(t *testing.T) {
	n := 4
	cell := []int{2, 2}
	runSample(t, n, cell)
}

func TestSample3(t *testing.T) {
	n := 5
	cell := []int{3, 3}
	runSample(t, n, cell)
}
