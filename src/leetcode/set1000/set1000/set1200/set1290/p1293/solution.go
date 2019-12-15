package p1293

const INF = 1 << 30

var dd = []int{-1, 0, 1, 0, -1}

func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	dist := make([][][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = make([]int, k+1)
			for u := 0; u <= k; u++ {
				dist[i][j][u] = -1
			}
		}
	}

	dist[0][0][0] = 0

	N := m * n * (k + 1)

	que := make([]int, N)
	var end int
	que[end] = 0
	end++
	var front int

	for front < end {
		cur := que[front]
		front++
		c := cur % (k + 1)
		b := (cur / (k + 1)) % n
		a := (cur / (k + 1)) / n

		for i := 0; i < 4; i++ {
			u, v := a+dd[i], b+dd[i+1]
			if u < 0 || u >= m || v < 0 || v >= n {
				continue
			}
			if grid[u][v] == 0 {
				if dist[u][v][c] < 0 {
					dist[u][v][c] = dist[a][b][c] + 1
					que[end] = u*n*(k+1) + v*(k+1) + c
					end++
				}
			} else if c+1 <= k {
				if dist[u][v][c+1] < 0 {
					dist[u][v][c+1] = dist[a][b][c] + 1
					que[end] = u*n*(k+1) + v*(k+1) + c + 1
					end++
				}
			}
		}
	}

	res := -1

	for i := 0; i <= k; i++ {
		if dist[m-1][n-1][i] < 0 {
			continue
		}
		if res < 0 || res > dist[m-1][n-1][i] {
			res = dist[m-1][n-1][i]
		}
	}

	return res
}

func shortestPath1(grid [][]int, k int) int {
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
