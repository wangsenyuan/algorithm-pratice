package p1463

func cherryPickup(grid [][]int) int {
	// r := len(grid)
	c := len(grid[0])

	dp := make([][][]int, 2)
	can := make([][][]bool, 2)
	for k := 0; k < 2; k++ {
		dp[k] = make([][]int, c)
		can[k] = make([][]bool, c)
		for i := 0; i < c; i++ {
			dp[k][i] = make([]int, c)
			can[k][i] = make([]bool, c)
		}
	}

	dp[0][0][c-1] = grid[0][0] + grid[0][c-1]
	can[0][0][c-1] = true

	update := func(r int, x int, y int, prev int) {
		if x < 0 || y < 0 || x >= c || y >= c {
			return
		}
		if x != y {
			dp[(r+1)&1][x][y] = max(dp[(r+1)&1][x][y], prev+grid[r+1][x]+grid[r+1][y])
		} else {
			dp[(r+1)&1][x][y] = max(dp[(r+1)&1][x][y], prev+grid[r+1][x])
		}
		can[(r+1)&1][x][y] = true
	}

	for r := 0; r < len(grid)-1; r++ {
		for i := 0; i < c; i++ {
			for j := 0; j < c; j++ {
				dp[(r+1)&1][i][j] = 0
			}
		}
		for i := 0; i < c; i++ {
			for j := 0; j < c; j++ {
				if !can[r&1][i][j] {
					continue
				}
				update(r, i-1, j-1, dp[r&1][i][j])
				update(r, i-1, j, dp[r&1][i][j])
				update(r, i-1, j+1, dp[r&1][i][j])
				update(r, i, j-1, dp[r&1][i][j])
				update(r, i, j, dp[r&1][i][j])
				update(r, i, j+1, dp[r&1][i][j])
				update(r, i+1, j-1, dp[r&1][i][j])
				update(r, i+1, j, dp[r&1][i][j])
				update(r, i+1, j+1, dp[r&1][i][j])
			}
		}
	}
	n := len(grid)
	var best int
	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {
			if can[(n-1)&1][i][j] {
				best = max(best, dp[(n-1)&1][i][j])
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
