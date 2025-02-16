package p3458

func maxSubstringLength(s string, k int) bool {
	if k == 0 {
		return true
	}
	// n := len(s)
	first := make([]int, 26)
	last := make([]int, 26)

	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}
	for i, x := range s {
		last[int(x-'a')] = i
		if first[int(x-'a')] < 0 {
			first[int(x-'a')] = i
		}
	}
	n := len(s)

	dp := make([]int, n+1)
	// dp[r] = l, 表示最短的必须和r分在一起的部分
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]
		x := int(s[i] - 'a')
		if last[x] == i {
			l := first[x]
			ok := true
			for j := i - 1; j >= l; j-- {
				y := int(s[j] - 'a')
				l = min(l, first[y])
				if last[y] > i {
					ok = false
					break
				}
			}
			if ok {
				dp[i+1] = max(dp[i], dp[l]+1)
				if i < n-1 && dp[i+1] >= k {
					return true
				}
				if k == 1 && i == n-1 && l == 0 {
					return false
				}
			}
		}
	}

	return dp[n] >= k
}
