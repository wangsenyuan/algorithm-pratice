package p3257

func countKConstraintSubstrings(s string, k int) int {

	n := len(s)
	fp := make([]int, n+1)
	pref := make([]int, n+1)
	cnt := make([]int, 2)
	for l, r := 0, 0; r < n; r++ {
		cnt[int(s[r]-'0')]++
		for cnt[0] > k && cnt[1] > k {
			cnt[int(s[l]-'0')]--
			l++
		}
		fp[r+1] = r + 1 - l + fp[r]
		pref[r+1] = pref[r] + l
	}
	dp := make([]int, n+1)
	next := make([]int, n+1)
	cnt[0] = 0
	cnt[1] = 0
	for l, r := n-1, n-1; l >= 0; l-- {
		cnt[int(s[l]-'0')]++
		for cnt[0] > k && cnt[1] > k {
			cnt[int(s[r]-'0')]--
			r--
		}
		next[l] = r
		// dp[l] = 区间l1...r 且 l1 < l 的部分
		dp[l] = (r-l+1)*l - (pref[r+1] - pref[l])
	}

	if next[0] == n-1 {
		return (n + 1) * n / 2
	}

	return fp[n] - dp[0]
}
