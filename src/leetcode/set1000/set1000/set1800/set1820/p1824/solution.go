package p1824

const INF = 1 << 30

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][1] = 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < 3; j++ {
			if dp[i][j] < INF {
				for k := 0; k < 3; k++ {
					if obstacles[i] == 0 || obstacles[i] != k+1 {
						if obstacles[i+1] == 0 || obstacles[i+1] != k+1 {
							if j == k {
								// no slide
								dp[i+1][k] = min(dp[i+1][k], dp[i][j])
							} else {
								// slide from j to k, then move
								dp[i+1][k] = min(dp[i+1][k], dp[i][j]+1)
							}
						}
					}
				}
			}
		}
	}
	res := INF
	for j := 0; j < 3; j++ {
		res = min(res, dp[n-1][j])
	}
	if res == INF {
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
