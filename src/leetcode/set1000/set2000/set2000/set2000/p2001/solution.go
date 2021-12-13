package p2001

func maxProduct(s string) int {
	n := len(s)
	// n <= 12
	N := 1 << n

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	pp := make([]int, N)

	buf := make([]byte, n)

	for mask := 0; mask < N; mask++ {
		var j int
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				buf[j] = s[i]
				j++
			}
		}
		pp[mask] = pl(buf[:j], dp)
	}

	var best int

	for mask := 1; mask < N; mask++ {
		x := pp[mask]
		y := pp[(N-1)^mask]

		tmp := x * y

		if tmp > best {
			best = tmp
		}
	}

	return best
}

func pl(s []byte, dp [][]int) int {
	if len(s) == 0 {
		return 0
	}

	n := len(s)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
		dp[i][i] = 1
	}

	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] {
				if i+1 <= j-1 {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
