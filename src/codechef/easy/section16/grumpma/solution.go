package main

func main() {

}

const MOD = 1000000007

func solve(n int, K int, M int, A []int) int {
	m := (K + M - 1) / M
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(x, idx int) int

	dfs = func(x, idx int) int {
		if idx >= n {
			return 0
		}
		if dp[x][idx] >= 0 {
			return dp[x][idx]
		}

		cnt := make(map[int]int)
		var res int
		for i := idx; i < n; i++ {
			a := A[i] % M
			if a == 1 {
				cnt[a] = modAdd(cnt[a], 1)
			} else if a == 0 {
				cnt[a] = modAdd(cnt[a], cnt[M - 1])
			} else {
				cnt[a] = modAdd(cnt[a], cnt[a - 1])
			}
			if a == 0 && cnt[a] > 0 && x > 0 {
				tmp := int64(dfs(x - 1, i + 1))
				tmp *= int64(cnt[a])
				tmp %= MOD
				res = modAdd(res, int(tmp))
			}
		}

		if x == 0 {
			res = cnt[K % M]
		} 
		dp[x][idx] = res
		return res
	}

	return dfs(K / M, 0)
}

func modAdd(a, b int) int {
	c := a + b
	if c >= MOD {
		c -= MOD
	}
	return c
}
