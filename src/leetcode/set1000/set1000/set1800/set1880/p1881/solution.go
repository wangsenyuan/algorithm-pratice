package p1881

const INF = 1<<31 - 1

func minimumXORSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	if n == 1 {
		return nums1[0] ^ nums2[0]
	}
	N := 1 << n
	dp := make([]int, N)

	for i := 1; i < N; i++ {
		dp[i] = INF
	}

	for state := 1; state < N; state++ {
		var i int
		for j := 0; j < n; j++ {
			if state&(1<<j) > 0 {
				i++
			}
		}
		for j := 0; j < n; j++ {
			if state&(1<<j) > 0 {
				dp[state] = min(dp[state], dp[state^(1<<j)]+(nums1[i-1]^nums2[j]))
			}
		}
	}
	return dp[N-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
