package p1473

const INF = 1 << 26

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// then always has answer
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, target+1)
			for k := 0; k <= target; k++ {
				dp[i][j][k] = INF
			}
		}
	}

	for i := 0; i < m; i++ {
		if houses[i] > 0 {
			// already has color
			if i == 0 {
				dp[i][houses[i]-1][1] = 0
			} else {
				for j := 0; j < n; j++ {
					for k := 0; k <= target; k++ {
						var nk = k
						if j != houses[i]-1 {
							nk++
						}
						if nk <= target {
							dp[i][houses[i]-1][nk] = min(dp[i][houses[i]-1][nk], dp[i-1][j][k])
						}
					}
				}
			}
		} else {
			//try assign it different colors
			for j := 0; j < n; j++ {
				if i == 0 {
					dp[i][j][1] = cost[i][j]
					continue
				}
				for jj := 0; jj < n; jj++ {
					for k := 0; k <= target; k++ {
						var nk = k
						if jj != j {
							nk++
						}
						if nk <= target {
							dp[i][j][nk] = min(dp[i][j][nk], dp[i-1][jj][k]+cost[i][j])
						}
					}
				}
			}
		}
	}

	var best = INF

	for j := 0; j < n; j++ {
		best = min(best, dp[m-1][j][target])
	}

	if best >= INF {
		return -1
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
