package p2370

func longestIdealString(s string, k int) int {

	n := len(s)

	dp := make([]int, 26)

	var res int

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		var tmp int
		for j := 0; j < 26; j++ {
			if abs(j-x) <= k {
				tmp = max(tmp, dp[j])
			}
		}
		dp[x] = max(dp[x], tmp+1)
		res = max(res, dp[x])
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
