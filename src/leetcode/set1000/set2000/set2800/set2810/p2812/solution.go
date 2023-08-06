package p2812

func canSplitArray(nums []int, m int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = nums[i]
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	get := func(l int, r int) int {
		res := sum[r]
		if l > 0 {
			res -= sum[l-1]
		}
		return res
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n+1)
		dp[i][1] = true
		dp[i][2] = true
	}

	for l := 3; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			if dp[i][l-1] && get(i, i+l-2) >= m {
				dp[i][l] = true
			}
			if dp[i+1][l-1] && get(i+1, i+l-1) >= m {
				dp[i][l] = true
			}
		}
	}
	return dp[0][n]
}
