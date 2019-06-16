package p995

func minKBitFlips(A []int, K int) int {
	n := len(A)
	dp := make([]int, n)

	var ans int
	var i = 0
	for i+K <= n {
		if i > 0 {
			dp[i] += dp[i-1]
		}
		r := A[i] + (dp[i] & 1)
		if r&1 == 0 {
			ans++
			//need to flip it
			dp[i]++
			if i+K < n {
				dp[i+K]--
			}
		}
		i++
	}

	for i < n {
		dp[i] += dp[i-1]
		r := A[i] + (dp[i] & 1)
		if r&1 == 0 {
			return -1
		}
		i++
	}

	return ans
}
