package p2301

func matchReplacement(s string, sub string, mappings [][]byte) bool {

	// 26 + 26 + 10
	table := make([][]bool, 62)
	for i := 0; i < 62; i++ {
		table[i] = make([]bool, 62)
	}

	for _, cur := range mappings {
		a, b := index(cur[0]), index(cur[1])
		table[a][b] = true
	}
	m := len(s)
	n := len(sub)

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	check := func(a, b byte) bool {
		return a == b || table[index(a)][index(b)]
	}

	for l := 1; l <= n; l++ {
		for i := l; i <= m; i++ {
			if l == 1 {
				// sub[:1] is a valid substring of s[:i]
				dp[i][l] = check(sub[0], s[i-1])
			} else if !dp[i][l] {
				dp[i][l] = i-1 >= 0 && dp[i-1][l-1] && check(sub[l-1], s[i-1])
			}
			if l == n && dp[i][l] {
				return true
			}
		}
	}

	return false
}

func index(x byte) int {
	if x >= 'a' && x <= 'z' {
		return int(x - 'a')
	}
	if x >= 'A' && x <= 'Z' {
		return 26 + int(x-'A')
	}
	return 52 + int(x-'0')
}
