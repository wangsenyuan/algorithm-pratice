package p2785

func numberOfWays(n int, x int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		// 迭代i，保证了i不会被重复选中
		v := pow(i, x)
		for s := n; s >= v; s-- {
			dp[s] = add(dp[s], dp[s-v])
		}
	}
	return dp[n]
}

const MOD = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n&1 == 1 {
			res = res * x
		}
		x = x * x
	}
	return res
}
