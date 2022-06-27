package p2401

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int64, n+1)
	}

	for _, price := range prices {
		h, w, p := price[0], price[1], price[2]
		dp[h][w] = max(dp[h][w], int64(p))
	}

	for h := 1; h <= m; h++ {
		for w := 1; w <= n; w++ {
			for h1 := 1; h1 < h; h1++ {
				h2 := h - h1
				dp[h][w] = max(dp[h][w], dp[h1][w]+dp[h2][w])
			}
			for w1 := 1; w1 < w; w1++ {
				w2 := w - w1
				dp[h][w] = max(dp[h][w], dp[h][w1]+dp[h][w2])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
