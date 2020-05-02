package p1429

const MOD = 1000000007

func numberWays(hats [][]int) int {
	n := len(hats)

	R := make([][]int, 40)

	for i := 0; i < 40; i++ {
		R[i] = make([]int, 0, n)
	}

	for i, hat := range hats {
		for j := 0; j < len(hat); j++ {
			x := hat[j] - 1
			R[x] = append(R[x], i)
		}
	}

	dp := make([][]int64, 40)

	for i := 0; i < 40; i++ {
		dp[i] = make([]int64, 1<<n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	N := 1 << n

	var dfs func(i int, set int) int64

	dfs = func(i int, set int) int64 {
		if set == N-1 {
			return 1
		}

		if i == 40 {
			return 0
		}
		if dp[i][set] >= 0 {
			return dp[i][set]
		}

		// skip i
		dp[i][set] = dfs(i+1, set)
		// try assign hat i to person
		for j := 0; j < len(R[i]); j++ {
			x := R[i][j]
			if set&(1<<x) > 0 {
				continue
			}
			dp[i][set] += dfs(i+1, set|(1<<x))
			dp[i][set] %= MOD
		}

		return dp[i][set]
	}

	res := dfs(0, 0)

	return int(res)
}

func numberWays1(hats [][]int) int {
	n := len(hats)

	var dfs func(i int, used int64) int64

	dfs = func(i int, used int64) int64 {
		if i == n {
			return 1
		}

		hat := hats[i]

		var res int64

		for j := 0; j < len(hat); j++ {
			x := hat[j]
			if used&(int64(1)<<x) > 0 {
				continue
			}
			res += dfs(i+1, used|(int64(1)<<x))
			res %= MOD
		}

		return res % MOD
	}

	return int(dfs(0, 0))
}
