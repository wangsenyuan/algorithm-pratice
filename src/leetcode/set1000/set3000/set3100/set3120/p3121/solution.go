package p3121

const oo = 1 << 60

func minimumOperations(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	// dp[0...9] 表示将当前列变成i后的最优值
	// dp[10] = 当前列变成任何值，且和前面不冲突的情况
	dp := make([]int, 11)
	for i := 0; i < 10; i++ {
		dp[i] = oo
	}
	dp[10] = 0
	cnt := make([]int, 10)
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			cnt[grid[i][j]]++
		}
		// 前面的最优值
		best := []int{-1, -1}

		for i := 0; i <= 10; i++ {
			if best[0] < 0 || dp[i] <= dp[best[0]] {
				best[1] = best[0]
				best[0] = i
			} else if best[1] < 0 || dp[i] <= dp[best[1]] {
				best[1] = i
			}
		}
		a, b := dp[best[0]], dp[best[1]]
		for x := 0; x < 10; x++ {
			dp[x] = oo
			if best[0] == x {
				dp[x] = b + n - cnt[x]
			} else {
				dp[x] = a + n - cnt[x]
			}
			cnt[x] = 0
		}
		dp[10] = a + n
	}

	res := dp[0]
	for i := 0; i <= 10; i++ {
		res = min(res, dp[i])
	}

	return res
}
