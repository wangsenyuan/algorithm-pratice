package p2209

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	// n = len(floor) <= 1000
	// dp[i][x] = 在第i位放置第x个carpet时所剩余的白色tile, 不超过 i + carpetLen
	// dp[i][x] = min(dp[j][x-1] + w[i-1] - w[j+carpetLen-1]) where j + carpetLen < i
	// dp[i][x] = w[i-1] + min(dp[j][x-1] - w[j+carpetLen-1])
	// let fp[j][x-1] = min(dp[j][x-1] - w[j+carpetLen]) from 0...j
	n := len(floor)
	W := make([]int, n)
	for i := 0; i < n; i++ {
		if floor[i] == '1' {
			W[i]++
		}
		if i > 0 {
			W[i] += W[i-1]
		}
	}

	dp := make([]int, n)

	fp := make([]int, n)

	for x := 1; x <= numCarpets; x++ {
		for i := n - 1; i >= 0; i-- {
			// put x-th carpet from position i
			if x == 1 {
				if i > 0 {
					dp[i] = W[i-1]
				} else {
					dp[i] = 0
				}
			} else {
				j := i - carpetLen
				// j + carpetLen <= i
				if j < 0 || i == 0 {
					dp[i] = 0
				} else {
					dp[i] = W[i-1] + fp[j]
				}
			}
			fp[i] = n
		}

		for i := 0; i < n; i++ {
			fp[i] = dp[i] - W[min(i+carpetLen, n)-1]
			if i > 0 {
				fp[i] = min(fp[i], fp[i-1])
			}
		}
	}

	res := n

	for i := 0; i < n; i++ {
		// put a carpet here
		tmp := dp[i] + W[n-1]
		if i+carpetLen <= n {
			tmp -= W[i+carpetLen-1]
		}
		res = min(res, tmp)
	}

	return res
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
