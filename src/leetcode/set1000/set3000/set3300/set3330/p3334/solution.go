package p3334

func maxScore(nums []int) int64 {
	type pair struct {
		first  int
		second int
	}

	n := len(nums)
	dp := make([]pair, n+1)
	// first is for lcm, second is for gcd
	dp[0] = pair{1, 0}
	for i, num := range nums {
		g := gcd(dp[i].second, num)
		l := lcm(dp[i].first, num)
		dp[i+1] = pair{l, g}
	}

	res := dp[n].first * dp[n].second
	suf := pair{1, 0}

	for i := n - 1; i >= 0; i-- {
		// 如果删除i
		x := gcd(suf.second, dp[i].second)
		y := lcm(suf.first, dp[i].first)
		res = max(res, x*y)
		g := gcd(suf.second, nums[i])
		l := lcm(suf.first, nums[i])
		suf = pair{l, g}
	}

	return int64(res)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}
