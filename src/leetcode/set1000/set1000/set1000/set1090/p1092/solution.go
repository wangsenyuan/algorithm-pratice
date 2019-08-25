package p1092

import "bytes"

const INF = 1 << 20

func shortestCommonSupersequence(str1 string, str2 string) string {
	m := len(str1)
	n := len(str2)

	// dp[i][j] = len of super-seq that contains s1[:i] & s2[:j]
	// dp[0][j] = j+1
	// dp[i][0] = i + 1
	// dp[i][j] = if s[i] == s[j] then dp[i][j] = dp[i-1][j-1] + 1
	// else dp[i][j] = dp[i-1][j-1] + 2
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	var buf bytes.Buffer

	a := m
	b := n

	for a > 0 && b > 0 {
		if str1[a-1] == str2[b-1] {
			buf.WriteByte(str1[a-1])
			a--
			b--
		} else if dp[a-1][b] < dp[a][b-1] {
			buf.WriteByte(str1[a-1])
			a--
		} else {
			buf.WriteByte(str2[b-1])
			b--
		}
	}

	for a > 0 {
		buf.WriteByte(str1[a-1])
		a--
	}

	for b > 0 {
		buf.WriteByte(str2[b-1])
		b--
	}

	res := buf.Bytes()

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
