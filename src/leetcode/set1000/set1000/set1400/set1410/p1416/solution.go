package p1416

const MOD = 1000000007

func numberOfArrays(s string, k int) int {

	n := len(s)
	dp := make([]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}

	var dfs func(i int) int64

	dfs = func(i int) int64 {
		if i == n {
			return 1
		}
		if dp[i] >= 0 {
			return dp[i]
		}
		if s[i] == '0' {
			return 0
		}

		dp[i] = 0

		var num int

		for j := i; j < n && num < k; j++ {
			num = num*10 + int(s[j]-'0')
			if num > k {
				break
			}
			dp[i] += dfs(j + 1)
			if dp[i] >= MOD {
				dp[i] -= MOD
			}
		}

		return dp[i]
	}

	return int(dfs(0))

}
