package p3494

const inf = 1 << 60

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)

	dp := make([]int, n)

	for j := range m {
		var sum int
		for i := range n {
			sum = max(sum, dp[i]) + skill[i]*mana[j]
		}
		dp[n-1] = sum
		for i := n - 2; i >= 0; i-- {
			sum -= skill[i+1] * mana[j]
			dp[i] = sum
		}
	}

	return int64(dp[n-1])
}
