package p3286

var dd = []int{-1, 0, 1, 0, -1}

func findSafeWalk(grid [][]int, health int) bool {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = health - grid[0][0]

	que := make([]int, 2*n*m)
	head, tail := n*m, n*m

	que[head] = 0
	head++

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++

		if dp[r][c] == 0 {
			continue
		}
		for i := 0; i < 4; i++ {
			u, v := r+dd[i], c+dd[i+1]
			if u < 0 || u == n || v < 0 || v == m || dp[u][v] >= 0 {
				continue
			}
			if grid[u][v] == 1 {
				dp[u][v] = dp[r][c] - 1
				que[head] = u*m + v
				head++
			} else {
				dp[u][v] = dp[r][c]
				tail--
				que[tail] = u*m + v
			}
		}
	}

	return dp[n-1][m-1] > 0
}
