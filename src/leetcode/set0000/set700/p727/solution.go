package p727

func minWindow(S string, T string) string {
	n, m := len(S), len(T)

	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)

	for i := 0; i < n; i++ {
		if S[i] == T[0] {
			dp[0][i] = i
		} else {
			dp[0][i] = -1
		}
	}

	for j := 1; j < m; j++ {
		for k := 0; k < n; k++ {
			dp[j&1][k] = -1
		}
		last := -1
		for i := 0; i < n; i++ {
			if last >= 0 && S[i] == T[j] {
				dp[j&1][i] = last
			} else if dp[1-(j&1)][i] >= 0 {
				last = dp[1-(j&1)][i]
			}
		}
	}

	start, end := 0, n
	for e := 0; e < n; e++ {
		s := dp[1-(m&1)][e]
		if s >= 0 && e-s < end-start {
			start = s
			end = e
		}
	}
	if end < n {
		return S[start : end+1]
	}
	return ""
}
