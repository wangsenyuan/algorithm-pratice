package p2876

func maximumTripletValue(nums []int) int64 {
	// dp[i][0] = 到i为止nums[a] - nums[b] 的最小的值 a < b
	// dp[i][1] = 到i为止nums[a] - nums[b] 的最大值 a < b
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 1 << 30
	dp[0][1] = -(1 << 30)
	// dp[0] is not valid
	minAt := 0
	maxAt := 0

	for i := 1; i < n; i++ {
		dp[i][0] = min(dp[i-1][0], nums[minAt]-nums[i])
		dp[i][1] = max(dp[i-1][1], nums[maxAt]-nums[i])
		if nums[minAt] > nums[i] {
			minAt = i
		}
		if nums[maxAt] < nums[i] {
			maxAt = i
		}
	}
	var best int64
	for k := 2; k < n; k++ {
		if nums[k] < 0 {
			best = max(best, int64(nums[k])*int64(dp[k-1][0]))
		} else {
			best = max(best, int64(nums[k])*int64(dp[k-1][1]))
		}
	}

	return best
}

type Number interface {
	int | int64
}

func max[T Number](a T, b T) T {
	if a >= b {
		return a
	}
	return b
}

func min[T Number](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}
