package p2767

func minimumBeautifulSubstrings(s string) int {
	n := len(s)
	// n <= 15
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}
		if s[i] == '0' {
			return n + 1
		}
		if dp[i] < 0 {
			dp[i] = n + 1
			var cur int
			for j := i; j < n; j++ {
				cur = cur*2 + int(s[j]-'0')
				if check(cur) {
					dp[i] = min(dp[i], dfs(j+1)+1)
				}
			}
		}

		return dp[i]
	}

	res := dfs(0)
	if res <= n {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func check(num int) bool {
	for num > 0 && num%5 == 0 {
		num /= 5
	}
	return num == 1
}
