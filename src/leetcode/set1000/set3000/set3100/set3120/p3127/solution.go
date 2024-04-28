package p3127

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][]int, zero+1)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][]int, one+1)
		for j := 0; j <= one; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 1; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 1; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}

	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			// 第i+j位放置0，前面放置0/1
			dp[i][j][0] = add(dp[i-1][j][0], dp[i-1][j][1])
			if i > limit {
				// 最多连续使用limit个0，如果在第i-limit-1+j位为1，这是一个bad的序列
				dp[i][j][0] = sub(dp[i][j][0], dp[i-limit-1][j][1])
			}
			dp[i][j][1] = add(dp[i][j-1][0], dp[i][j-1][1])
			if j > limit {
				dp[i][j][1] = sub(dp[i][j][1], dp[i][j-limit-1][0])
			}
		}
	}

	return add(dp[zero][one][0], dp[zero][one][1])
}
