package p2087

func countPyramids(grid [][]int) int {
	res := count(grid)
	for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
		grid[i], grid[j] = grid[j], grid[i]
	}

	res += count(grid)

	return res
}

func count(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var res int
	for r := m - 1; r >= 0; r-- {
		// dp[r][c] = max height of the pyramid with r, c as apex
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				dp[r][c] = 0
				continue
			}
			if r == m-1 || c == 0 || c == n-1 {
				dp[r][c] = 1
				continue
			}
			dp[r][c] = min(dp[r+1][c], min(dp[r+1][c-1], dp[r+1][c+1])) + 1
			res += dp[r][c] - 1
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
