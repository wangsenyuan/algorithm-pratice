package domino

func domino(n int, m int, broken [][]int) int {
	bad := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		bad[i] = make([]bool, m)
	}

	for _, cell := range broken {
		u, v := cell[0], cell[1]
		bad[u][v] = true
	}

	for j := 0; j < m; j++ {
		bad[n][j] = true
	}

	M := 1 << m

	dp := make([][][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, M)
			for k := 0; k < M; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var cur int
	for j := 0; j < m; j++ {
		if bad[0][j] {
			cur |= 1 << (m - 1 - j)
		}
	}
	var best int
	dp[0][0][cur] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ii := i
			jj := j + 1
			if jj == m {
				ii++
				jj = 0
			}
			for state := 0; state < M; state++ {
				if dp[i][j][state] < 0 {
					continue
				}
				best = max(best, dp[i][j][state])
				next := (state << 1) & (M - 1)

				if bad[i+1][j] {
					next |= 1
				}

				dp[ii][jj][next] = max(dp[ii][jj][next], dp[i][j][state])

				if j+1 < m {
					if state&(1<<(m-1)) == 0 && next&(1<<(m-1)) == 0 {
						dp[ii][jj][next|(1<<(m-1))] = max(dp[ii][jj][next|(1<<(m-1))], dp[i][j][state]+1)
					}
				}

				if state&(1<<(m-1)) == 0 && !bad[i+1][j] {
					dp[ii][jj][next|1] = max(dp[ii][jj][next|1], dp[i][j][state]+1)
				}
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
