package p1293

const INF = 1 << 30

var dd = []int{-1, 0, 1, 0, -1}

func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]bool, m)

	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	dp := make([][][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k+1)
			for u := 0; u <= k; u++ {
				dp[i][j][u] = INF
			}
		}
	}

	var dfs func(x int, y int, k int, cnt int) int

	best := INF

	dfs = func(x, y, k, cnt int) int {
		if cnt > best {
			return best
		}

		if x == m-1 && y == n-1 {
			best = cnt
			return cnt
		}

		if dp[x][y][k] < INF {
			return dp[x][y][k]
		}

		var res = INF

		for i := 0; i < 4; i++ {
			u, v := x+dd[i], y+dd[i+1]
			if u < 0 || u >= m || v < 0 || v >= n || vis[u][v] {
				continue
			}
			vis[u][v] = true

			if grid[u][v] == 0 {
				res = min(res, dfs(u, v, k, cnt+1))
			} else if k > 0 {
				res = min(res, dfs(u, v, k-1, cnt+1))
			}

			vis[u][v] = false
		}

		dp[x][y][k] = res

		return res
	}

	vis[0][0] = true

	res := dfs(0, 0, k, 0)

	if res >= INF {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
