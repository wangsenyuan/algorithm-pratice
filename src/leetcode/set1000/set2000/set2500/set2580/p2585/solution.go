package p2585

func waysToReachTarget(target int, types [][]int) int {
	// target <= 1000
	n := len(types)
	// n <= 50
	dp := make([]int, target+1)
	dp[0] = 1

	fp := make([]int, target+1)

	for i := 0; i < n; i++ {
		cur := types[i]
		cnt, mark := cur[0], cur[1]
		if cnt*mark > target {
			cnt = target / mark
		}
		for j := 0; j <= target; j++ {
			fp[j] = 0
		}
		for j := cnt; j > 0; j-- {
			for x := target; x-j*mark >= 0; x-- {
				fp[x] = modAdd(fp[x], dp[x-j*mark])
			}
		}
		for j := 0; j <= target; j++ {
			dp[j] = modAdd(dp[j], fp[j])
		}
	}

	return dp[target]
}

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}
