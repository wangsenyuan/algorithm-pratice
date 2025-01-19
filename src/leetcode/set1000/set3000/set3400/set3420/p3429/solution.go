package p3429

import "slices"

const inf = 1 << 60

func minCost(n int, cost [][]int) int64 {
	// 从两头开始
	dp := make([]int, 9)
	// dp[i][j]表示两头的房子
	for i := range 9 {
		dp[i] = inf
	}
	fp := make([]int, 9)
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if i == 0 {
			for u := 0; u < 3; u++ {
				for v := 0; v < 3; v++ {
					if u == v {
						continue
					}
					dp[u*3+v] = cost[i][u] + cost[j][v]
				}
			}
		} else {
			if i < j {
				for j := range 9 {
					fp[j] = inf
				}
				for u := 0; u < 3; u++ {
					for v := 0; v < 3; v++ {
						if u == v {
							continue
						}
						for a := 0; a < 3; a++ {
							if a == u {
								continue
							}
							for b := 0; b < 3; b++ {
								if a == b || b == v {
									continue
								}
								fp[a*3+b] = min(fp[a*3+b], dp[u*3+v]+cost[i][a]+cost[j][b])
							}
						}
					}
				}
				copy(dp, fp)
			} else {
				ans := inf
				for u := 0; u < 3; u++ {
					for v := 0; v < 3; v++ {
						if u == v {
							continue
						}
						w := 3 - u - v
						ans = min(ans, dp[u*3+v]+cost[i][w])
					}
				}
				return int64(ans)
			}
		}
	}

	return int64(slices.Min(dp))
}
