package p1960

const INF = 1 << 29

func minSpaceWastedKResizing(nums []int, k int) int {
	//dp[sz][i]为进入i时为sz，浪费的空间数
	// sz >= max(arr[i...]) 为 cnt * sz - sum[i...]
	// 所以sz没关系，只需要记录从i开始保证能满足所有的要求即可
	n := len(nums)
	// 全部的sum <= 10^^8
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = nums[i] + sum[i]
	}
	// dp[i][j] 为到i为止调整j次的最小值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = INF
		}
	}
	for i := 0; i < n; i++ {
		//如果我们在i处调整一次或不调整
		tmp := nums[i]
		for j := i; j >= 0; j-- {
			tmp = max(tmp, nums[j])
			need := (i-j+1)*tmp - (sum[i+1] - sum[j])
			if j > 0 {
				for u := k; u > 0; u-- {
					dp[i][u] = min(dp[i][u], need+dp[j-1][u-1])
				}
			} else {
				dp[i][0] = (i+1)*tmp - sum[i+1]
			}
		}
	}

	return dp[n-1][k]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
