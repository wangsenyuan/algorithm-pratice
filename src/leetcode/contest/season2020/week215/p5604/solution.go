package p5604

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	if m < n {
		m, n = n, m
	}
	M := 1
	for i := 0; i < n; i++ {
		M *= 3
	}
	MM := M / 3
	dp := make([][][]int, M)
	fp := make([][][]int, M)
	for i := 0; i < M; i++ {
		dp[i] = make([][]int, introvertsCount+1)
		fp[i] = make([][]int, introvertsCount+1)
		for j := 0; j <= introvertsCount; j++ {
			dp[i][j] = make([]int, extrovertsCount+1)
			fp[i][j] = make([]int, extrovertsCount+1)
			for k := 0; k <= extrovertsCount; k++ {
				dp[i][j][k] = -1000000
				fp[i][j][k] = -1000000
			}
		}
	}
	dp[0][0][0] = 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			for mask := 0; mask < M; mask++ {
				for x := 0; x <= introvertsCount; x++ {
					for y := 0; y <= extrovertsCount; y++ {
						if dp[mask][x][y] < 0 {
							continue
						}
						if x < introvertsCount {
							d := 120
							if mask/MM == 1 {
								d -= 60
							} else if mask/MM == 2 {
								d -= 10
							}
							if c > 0 && mask%3 == 1 {
								d -= 60
							} else if c > 0 && mask%3 == 2 {
								d -= 10
							}
							next := mask*3%M + 1
							max(&fp[next][x+1][y], dp[mask][x][y]+d)
						}
						if y < extrovertsCount {
							d := 40
							if mask/MM == 1 {
								d -= 10
							} else if mask/MM == 2 {
								d += 40
							}
							if c > 0 && mask%3 == 1 {
								d -= 10
							} else if c > 0 && mask%3 == 2 {
								d += 40
							}
							next := mask*3%M + 2
							max(&fp[next][x][y+1], dp[mask][x][y]+d)
						}
						next := mask * 3 % M
						max(&fp[next][x][y], dp[mask][x][y])
					}
				}
			}

			for mask := 0; mask < M; mask++ {
				for i := 0; i <= introvertsCount; i++ {
					for j := 0; j <= extrovertsCount; j++ {
						dp[mask][i][j] = fp[mask][i][j]
						fp[mask][i][j] = -1000000
					}
				}
			}
		}
	}

	var best int

	for mask := 0; mask < M; mask++ {
		for i := 0; i <= introvertsCount; i++ {
			for j := 0; j <= extrovertsCount; j++ {
				max(&best, dp[mask][i][j])
			}
		}
	}
	return best
}

func max(a *int, b int) {
	if *a < b {
		*a = b
	}
}
