package p2901

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	last := make([]int, 2)
	last[0] = -1
	last[1] = -1

	dp := make([]int, n)
	// dp[0] = 最长的以0 = group[i] 的序列长度
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = -1
	}
	for i := 0; i < n; i++ {
		x := groups[i]
		if last[1-x] < 0 {
			dp[i] = 1
		} else {
			dp[i] = 1 + dp[last[1-x]]
			prev[i] = last[1-x]
		}
		last[x] = i
	}
	x := 0
	if last[0] < 0 || last[1] >= 0 && dp[last[1]] > dp[last[0]] {
		x = 1
	}

	ln := dp[last[x]]
	pos := last[x]
	ans := make([]string, ln)

	for ln > 0 {
		ans[ln-1] = words[pos]
		pos = prev[pos]
		ln--
	}

	return ans
}
