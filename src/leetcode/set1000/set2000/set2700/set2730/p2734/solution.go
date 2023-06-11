package p2734

func minCost(nums []int, x int) int64 {
	// 假设要手机巧克力i
	// 移动到某个位置j后的成本为 abs(j - i) * x + nums[j]
	n := len(nums)
	dp := make([]int, n)
	copy(dp, nums)
	var best int64 = 1 << 62

	for r := 0; r < n; r++ {
		// i 处的可以移动到 (i + r) % n处, 那么需要知道这个区间里面的最小值
		tmp := int64(r) * int64(x)
		for i := 0; i < n; i++ {
			tmp += int64(dp[i])
		}
		if tmp < best {
			best = tmp
		}
		// 移动一次
		dp0 := dp[0]
		for i := 0; i+1 < n; i++ {
			dp[i] = min(nums[i], dp[i+1])
		}
		dp[n-1] = min(dp0, nums[n-1])
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
