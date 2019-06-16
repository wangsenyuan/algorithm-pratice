package p994

const INF = 1 << 20

var dd = []int{-1, 0, 1, 0, -1}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	que := make([]int, m*n)
	var front, end int
	dp := make([]int, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i*n+j] = -1
			if grid[i][j] == 2 {
				que[end] = i*n + j
				dp[i*n+j] = 0
				end++
			}
		}
	}

	for front < end {
		cur := que[front]
		front++
		x, y := cur/n, cur%n
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n {
				if dp[u*n+v] < 0 && grid[u][v] == 1 {
					que[end] = u*n + v
					dp[u*n+v] = dp[cur] + 1
					end++
				}
			}
		}
	}
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if dp[i*n+j] < 0 {
					return -1
				}
				ans = max(ans, dp[i*n+j])
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
