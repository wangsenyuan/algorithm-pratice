package p1596

import "math"

func connectTwoGroups(cost [][]int) int {
	m := len(cost)
	n := len(cost[0])
	N := 1 << n
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = math.MaxInt32 >> 1
		}
	}

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, N)
		for state := 0; state < N; state++ {
			var tmp int
			for j := 0; j < n; j++ {
				if state&(1<<j) > 0 {
					tmp += cost[i][j]
				}
			}
			grid[i][state] = tmp
		}
	}

	for state := 1; state < N; state++ {
		dp[0][state] = grid[0][state]
	}

	for i := 1; i < m; i++ {
		for state := 1; state < N; state++ {
			for j := 0; j < n; j++ {
				dp[i][state|1<<j] = min(dp[i][state|1<<j], dp[i-1][state]+cost[i][j])
			}
			rest := N - 1 - state

			for j := rest; j > 0; j = rest & (j - 1) {
				dp[i][j|state] = min(dp[i][j|state], dp[i-1][state]+grid[i][j])
			}
		}
	}
	return dp[m-1][N-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
