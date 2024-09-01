package p3276

func maxScore(grid [][]int) int {
	pos := make([][]int, 101)
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := grid[i][j]
			if len(pos[x]) == 0 || pos[x][len(pos[x])-1] != i {
				pos[x] = append(pos[x], i)
			}
		}
	}

	dp := make([][]int, 101)
	for i := 0; i < 101; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(x int, state int) (res int)

	dfs = func(x int, state int) (res int) {
		if x == 0 {
			return 0
		}
		if dp[x][state] >= 0 {
			return dp[x][state]
		}

		defer func() {
			dp[x][state] = res
		}()
		res = dfs(x-1, state)

		for _, r := range pos[x] {
			if (state>>r)&1 == 1 {
				continue
			}
			res = max(res, x+dfs(x-1, state|(1<<r)))
		}

		return res
	}

	return dfs(100, 0)
}
