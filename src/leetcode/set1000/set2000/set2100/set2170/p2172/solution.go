package p2172

import "math/bits"

func maximumANDSum(nums []int, numSlots int) int {
	// slots <= 9
	// n <= 2 * slots <= 18
	// nums[i] <= 15
	m := 2 * numSlots
	T := 1 << m
	dp := make([]int, T)
	for i := 0; i < len(dp); i++ {
		dp[i] = -(1 << 30)
	}
	dp[0] = 0

	var best int

	for state := 0; state < T; state++ {
		best = max(best, dp[state])
		c := bits.OnesCount(uint(state))

		if c >= len(nums) {
			continue
		}

		for i := 0; i < m; i++ {
			if (state>>i)&1 == 0 {
				next := state ^ (1 << i)
				dp[next] = max(dp[next], dp[state]+nums[c]&(i/2+1))
			}
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
