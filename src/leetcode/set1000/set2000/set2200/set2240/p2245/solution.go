package p2245

func maxTrailingZeros(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	fp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		fp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := grid[i][j]
			for num%2 == 0 {
				dp[i][j]++
				num /= 2
			}
			for num%5 == 0 {
				fp[i][j]++
				num /= 5
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
				fp[i][j] += fp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
				fp[i][j] += fp[i][j-1]
			}
			if i > 0 && j > 0 {
				dp[i][j] -= dp[i-1][j-1]
				fp[i][j] -= fp[i-1][j-1]
			}
		}
	}

	getRow := func(p [][]int, r, c int) int {
		res := p[r][c]
		if r > 0 {
			res -= p[r-1][c]
		}
		return res
	}

	getCol := func(p [][]int, r, c int) int {
		res := p[r][c]
		if c > 0 {
			res -= p[r][c-1]
		}
		return res
	}

	getAt := func(p [][]int, r, c int) int {
		res := p[r][c]
		if r > 0 {
			res -= p[r-1][c]
		}
		if c > 0 {
			res -= p[r][c-1]
		}
		if r > 0 && c > 0 {
			res += p[r-1][c-1]
		}
		return res
	}

	topLeft := func(x, y int) int {
		a := getCol(dp, x, y) + getRow(dp, x, y) - getAt(dp, x, y)
		b := getCol(fp, x, y) + getRow(fp, x, y) - getAt(fp, x, y)
		return min(a, b)
	}

	topRight := func(x, y int) int {
		a := getCol(dp, x, y) + getRow(dp, x, n-1) - getRow(dp, x, y)
		b := getCol(fp, x, y) + getRow(fp, x, n-1) - getRow(fp, x, y)
		return min(a, b)
	}

	bottomLeft := func(x, y int) int {
		a := getCol(dp, m-1, y) - getCol(dp, x, y) + getRow(dp, x, y)
		b := getCol(fp, m-1, y) - getCol(fp, x, y) + getRow(fp, x, y)
		return min(a, b)
	}

	bottomRight := func(x, y int) int {
		a := getCol(dp, m-1, y) - getCol(dp, x, y) + getRow(dp, x, n-1) - getRow(dp, x, y) + getAt(dp, x, y)
		b := getCol(fp, m-1, y) - getCol(fp, x, y) + getRow(fp, x, n-1) - getRow(fp, x, y) + getAt(fp, x, y)
		return min(a, b)
	}

	var best int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			best = max(best, topLeft(i, j))
			best = max(best, topRight(i, j))
			best = max(best, bottomLeft(i, j))
			best = max(best, bottomRight(i, j))
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
