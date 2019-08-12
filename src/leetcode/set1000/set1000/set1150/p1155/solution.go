package p1155

const MOD = 1e9 + 7

func numRollsToTarget(d int, f int, target int) int {
	dp := make([]int64, target+1)
	fp := make([]int64, target+1)

	dp[0] = 1

	for i := 1; i <= d; i++ {
		for x := 0; x <= target; x++ {
			fp[x] = 0
		}
		for x := 0; x < target; x++ {
			fp[x+1] += dp[x]
			if x+f+1 <= target {
				fp[x+f+1] -= dp[x]
			}
		}
		for x := 1; x <= target; x++ {
			fp[x] += fp[x-1]
			if fp[x] >= MOD {
				fp[x] -= MOD
			}
			if fp[x] >= MOD {
				fp[x] -= MOD
			}
			if fp[x] < 0 {
				fp[x] += MOD
			}
		}
		dp, fp = fp, dp

	}
	return int(dp[target])
}

func numRollsToTarget2(d int, f int, target int) int {
	dp := make([]int64, target+1)
	fp := make([]int64, target+1)
	dp[0] = 1

	for i := 1; i <= d; i++ {
		for x := 0; x <= target; x++ {
			fp[x] = 0
		}
		for x := target; x > 0; x-- {
			for y := f; y > 0; y-- {
				if x >= y {
					fp[x] += dp[x-y]
					if fp[x] >= MOD {
						fp[x] -= MOD
					}
				}
			}
		}
		dp, fp = fp, dp
	}
	return int(dp[target])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func numRollsToTarget1(d int, f int, target int) int {
	dp := make([][]int64, d+1)
	for i := 0; i <= d; i++ {
		dp[i] = make([]int64, target+1)
	}

	dp[0][0] = 1

	for i := 1; i <= d; i++ {
		for x := target; x > 0; x-- {
			for y := min(x, f); y > 0; y-- {
				dp[i][x] += dp[i-1][x-y]
				if dp[i][x] >= MOD {
					dp[i][x] -= MOD
				}
			}
		}
	}
	return int(dp[d][target])
}
