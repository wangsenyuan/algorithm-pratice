package p1529

const INF = 1 << 30

func getLengthOfOptimalCompression(s string, K int) int {
	n := len(s)
	dp := make([][][]int, 26)
	fp := make([][][]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = make([][]int, n+1)
		fp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, K+1)
			fp[i][j] = make([]int, K+1)
			for k := 0; k <= K; k++ {
				dp[i][j][k] = INF
				fp[i][j][k] = INF
			}
		}
		dp[i][0][0] = 0
	}
	//dp[ind(s[0])][1][0] = 1
	//dp[ind(s[0])][1][1] = 0

	// dp[x][l][k] = 以l ge x结尾，且删除掉k个字符的最少长度
	for u := 0; u < n; u++ {
		v := ind(s[u])

		for i := 0; i < 26; i++ {
			for j := 0; j <= u+1; j++ {
				for k := 0; k <= K && k <= (u+1); k++ {
					fp[i][j][k] = INF
				}
			}
		}
		for i := 0; i < 26; i++ {
			for j := 0; j <= u; j++ {
				for k := 0; k <= K && k <= (u+1); k++ {
					if dp[i][j][k] >= INF {
						// invalid
						continue
					}
					if k < K {
						// remove this one
						fp[i][j][k+1] = min(fp[i][j][k+1], dp[i][j][k])
					}
					if i == v {
						fp[i][j+1][k] = min(fp[i][j+1][k], dp[i][j][k]+calcDiff(j, j+1))
					} else {
						fp[v][1][k] = min(fp[v][1][k], dp[i][j][k]+1)
					}
				}
			}
		}
		dp, fp = fp, dp
	}

	var res = INF

	for i := 0; i < 26; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= K; k++ {
				res = min(res, dp[i][j][k])
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func calcDiff(x, y int) int {
	return numLength(y) - numLength(x)
}

func numLength(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	var res int
	for num > 0 {
		res++
		num /= 10
	}
	return res + 1
}

func ind(x byte) int {
	return int(x - 'a')
}
