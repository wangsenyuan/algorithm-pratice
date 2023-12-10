package p2959

func removeAlmostEqualCharacters(word string) int {
	n := len(word)
	// dp[1] = 当前位改变时的最小值
	// dp[0] = 当前不变时的最小值
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	for i := 1; i < n; i++ {
		// similar
		if dist(word[i], word[i-1]) <= 1 {
			// need to change
			// 保持当前位不变
			x := dp[1]
			// 改变当前位
			y := min(dp[0], dp[1]) + 1
			dp[0], dp[1] = x, y
		} else {
			// 不变也ok
			x := min(dp[0], dp[1])
			y := min(dp[0], dp[1]) + 1
			dp[0], dp[1] = x, y
		}
	}
	return min(dp[0], dp[1])
}

func dist(a, b byte) int {
	if a < b {
		a, b = b, a
	}
	return int(a-'a') - int(b-'a')
}
