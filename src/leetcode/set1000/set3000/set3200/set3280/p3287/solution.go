package p3287

func maxValue(nums []int, k int) int {
	n := len(nums)
	T := 1 << 7

	dp := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]bool, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]bool, T)
		}
	}

	dp[0][0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for s := 0; s < T; s++ {
				dp[i+1][j][s] = dp[i][j][s]
			}
		}
		for j := 0; j < k; j++ {
			for s := 0; s < T; s++ {
				if dp[i][j][s] {
					dp[i+1][j+1][s|nums[i]] = true
				}
			}
		}
	}

	fp := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		fp[i] = make([][]bool, k+1)
		for j := 0; j <= k; j++ {
			fp[i][j] = make([]bool, T)
		}
	}

	fp[n][0][0] = true

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= k; j++ {
			for s := 0; s < T; s++ {
				fp[i][j][s] = fp[i+1][j][s]
			}
		}

		for j := 0; j < k; j++ {
			for s := 0; s < T; s++ {
				if fp[i+1][j][s] {
					fp[i][j+1][s|nums[i]] = true
				}
			}
		}
	}

	var res int
	for i := k - 1; i+k <= n; i++ {
		for x := 0; x < T; x++ {
			for y := 0; y < T; y++ {
				if dp[i+1][k][x] && fp[i+1][k][y] {
					res = max(res, x^y)
				}
			}
		}
	}

	return res
}
