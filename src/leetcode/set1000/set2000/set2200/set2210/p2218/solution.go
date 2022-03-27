package p2218

const MAX_X = 2010
const N_INF = -(1 << 30)

func maxValueOfCoins(piles [][]int, k int) int {
	// n := len(piles)
	// n <= 1000, k <= 2000
	// 没有办法每个组合都尝试一遍
	// sum(cnt[i]) = k
	// max(sum(pref[i]))
	// cost[i][j] = j
	// value[i][j] = pref_sum[j]
	dp := make([]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = N_INF
	}


	dp[0] = 0
	for _, pile := range piles {
		fp := make([]int, k+1)

		var sum int
		for i := 0; i < len(pile); i++ {
			sum += pile[i]
			for j := 0; j+i < k; j++ {
				fp[j+i+1] = max(fp[j+i+1], dp[j]+sum)
			}
		}

		for i := 0; i <= k; i++ {
			dp[i] = max(dp[i], fp[i])
			if i > 0 {
				dp[i] = max(dp[i], dp[i-1])
			}
		}

	}

	return dp[k]
}

func maxOf(arr []int) int {
	var res int
	for _, num := range arr {
		res = max(res, num)
	}
	return res
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
