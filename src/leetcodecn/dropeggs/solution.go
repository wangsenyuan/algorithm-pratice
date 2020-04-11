package dropeggs

func superEggDrop(K int, N int) int {
	F := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		F[i] = make([]int, K+1)
	}

	for i := 1; i <= N; i++ {
		F[i][1] = i
	}

	for k := 1; k <= K; k++ {
		F[1][k] = 1
	}
	ans := 1
	for i := 2; i <= N; i++ {
		for k := 1; k <= K; k++ {
			F[i][k] = 1 + F[i-1][k] + F[i-1][k-1]
		}
		if F[i][K] >= N {
			ans = i
			break
		}
	}

	return ans
}

func superEggDrop2(K int, N int) int {
	dp := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = i
	}

	for k := 2; k <= K; k++ {
		dp2 := make([]int, N+1)
		var x = 1
		for n := 1; n <= N; n++ {

			for x < n && max(dp[x-1], dp2[n-x]) > max(dp[x], dp2[n-x-1]) {
				x++
			}
			dp2[n] = 1 + max(dp[x-1], dp2[n-x])
		}

		dp = dp2
	}

	return dp[N]
}

func superEggDrop1(K int, N int) int {
	mem := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		mem[i] = make([]int, N+1)
		for j := 0; j <= N; j++ {
			mem[i][j] = -1
		}
	}

	var dp func(k, n int) int

	dp = func(k, n int) int {
		if mem[k][n] >= 0 {
			return mem[k][n]
		}
		if n == 0 {
			return 0
		}

		if n == 1 {
			return 1
		}

		if k == 1 {
			return n
		}

		var left, right = 1, n

		for left < right {
			x := (left + right) / 2
			a := dp(k-1, x-1)
			b := dp(k, n-x)

			if a < b {
				left = x + 1
			} else {
				right = x
			}
		}
		// dp(k - 1, right - 1) >= dp(k, n - right)
		// dp(k - 1, right - 2) < dp(k, n - (right - 1))
		mem[k][n] = 1 + min(dp(k-1, right-1), dp(k, n-(right-1)))
		return mem[k][n]
	}

	return dp(K, N)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
