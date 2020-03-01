package p1368

var dd = []int{-1, 0, 1, 0, -1}

func minCost(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	N := m * n

	dp := make([][]int, m*n)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = m + n
		}
	}

	dp[m-1][n-1] = 0

	que := make([]int, 2*N)

	front, end := N, N

	que[end] = N - 1
	end++

	for front < end {
		cur := que[front]
		front++
		x, y := cur/n, cur%n
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n {
				var sameGroup bool
				if grid[u][v] == 1 && k == 3 {
					// u, v at left & point to right
					sameGroup = true
				} else if grid[u][v] == 2 && k == 1 {
					// u, v at right & point to left
					sameGroup = true
				} else if grid[u][v] == 3 && k == 0 {
					// u, v at top and point to down
					sameGroup = true
				} else if grid[u][v] == 4 && k == 2 {
					// u, v at down and point to up
					sameGroup = true
				}

				if sameGroup {
					if dp[u][v] > dp[x][y] {
						dp[u][v] = dp[x][y]
						front--
						que[front] = u*n + v
					}
				} else {
					if dp[u][v] > dp[x][y]+1 {
						dp[u][v] = dp[x][y] + 1
						que[end] = u*n + v
						end++
					}
				}
			}
		}
	}

	return dp[0][0]
}
