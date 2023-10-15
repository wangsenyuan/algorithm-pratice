package p2902

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	prev := make([]int, n)
	dp := make([]int, n)
	best := -1
	for j := 0; j < n; j++ {
		prev[j] = -1
		dp[j] = 1
		for i := 0; i < j; i++ {
			if groups[i] != groups[j] && distance(words[i], words[j]) == 1 {
				if dp[j] < dp[i]+1 {
					dp[j] = dp[i] + 1
					prev[j] = i
				}
			}
		}
		if best < 0 || dp[j] > dp[best] {
			best = j
		}
	}
	ans := make([]string, dp[best])
	it := dp[best]
	for it > 0 {
		ans[it-1] = words[best]
		best = prev[best]
		it--
	}

	return ans
}

func distance(a, b string) int {
	if len(a) != len(b) {
		return 2
	}
	var cnt int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
		if cnt > 1 {
			return 2
		}
	}
	return cnt
}
