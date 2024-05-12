package p3146

func minimumSubstringsInPartition(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	dp[0] = 0

	freq := make([]int, 26)
	var most int
	var cnt int
	add := func(x int, m int) bool {
		freq[x]++
		if freq[x] == 1 {
			cnt++
		}
		most = max(most, freq[x])

		return most*cnt == m
	}

	for i := 0; i < n; i++ {
		dp[i+1] = i + 1

		for j := 0; j < 26; j++ {
			freq[j] = 0
		}
		most = 0
		cnt = 0

		for j := i; j >= 0; j-- {
			if add(int(s[j]-'a'), i-j+1) {
				dp[i+1] = min(dp[i+1], dp[j]+1)
			}
		}
	}

	return dp[n]
}
