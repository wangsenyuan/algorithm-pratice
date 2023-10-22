package p2910

const inf = 1 << 30

func minimumChanges(s string, k int) int {
	n := len(s)

	check := func(l int, r int, d int) int {
		blck := (r - l + 1) / d
		var cnt int
		for i, j := 0, blck-1; i < j; i, j = i+1, j-1 {
			for u := 0; u < d; u++ {
				if s[l+i*d+u] != s[l+j*d+u] {
					cnt++
				}
			}
		}
		return cnt
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 0
		for j := i + 1; j < n; j++ {
			dp[i][j] = (j - i + 1) / 2
			// dp[i][j] = 将[i..j]做为一个子串后，变成半回文的最小次数
			ln := j - i + 1
			for d := 1; d*d <= ln; d++ {
				if ln%d == 0 {
					dp[i][j] = min(dp[i][j], check(i, j, d))
					if d != ln/d && d != 1 {
						dp[i][j] = min(dp[i][j], check(i, j, ln/d))
					}
				}
			}
		}
	}

	fp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		fp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			fp[i][j] = inf
		}
	}
	fp[0][0] = 0

	for i := 1; i <= n; i++ {
		// 长度至少位2
		for j := i - 2; j >= 0; j-- {
			for x := k; x > 0; x-- {
				fp[i][x] = min(fp[i][x], dp[j][i-1]+fp[j][x-1])
			}
		}
	}

	return fp[n][k]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
