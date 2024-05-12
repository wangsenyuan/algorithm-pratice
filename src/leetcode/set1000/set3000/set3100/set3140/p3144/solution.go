package p3144

const oo = 1 << 50

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	best := -oo
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = energy[i]
		if i+k < n {
			dp[i] += dp[i+k]
		}
		best = max(best, dp[i])
	}

	return best
}
