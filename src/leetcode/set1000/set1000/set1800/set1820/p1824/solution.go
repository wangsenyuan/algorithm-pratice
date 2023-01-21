package p1824

const INF = 1 << 30

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	// the min result when frog is at lan i
	dp := make([]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = INF
	}
	dp[1] = 0
	fp := make([]int, 3)

	for i := 0; i < n-1; i++ {
		for j := 0; j < 3; j++ {
			fp[j] = INF
		}
		for j := 0; j < 3; j++ {
			if dp[j] < INF {
				for k := 0; k < 3; k++ {
					if obstacles[i] == 0 || obstacles[i] != k+1 {
						if obstacles[i+1] == 0 || obstacles[i+1] != k+1 {
							x := dp[j]
							if j != k {
								// no slide
								x++
							}
							fp[k] = min(fp[k], x)
						}
					}
				}
			}
		}
		copy(dp, fp)
	}
	res := INF
	for j := 0; j < 3; j++ {
		res = min(res, dp[j])
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
