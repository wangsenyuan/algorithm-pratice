package p3474

func generateString(str1 string, str2 string) string {
	lps := kmp(str2)

	n := len(str1)
	m := len(str2)

	ok := make([]bool, m+1)
	for j := m; j > 0; j = lps[j-1] {
		ok[j] = true
	}

	buf := make([]byte, n+m-1)
	for i := range len(buf) {
		buf[i] = '-'
	}
	var p int
	for i := range n {
		if str1[i] == 'T' {
			if p <= i {
				copy(buf[i:], str2)
			} else {
				if !ok[p-i] {
					return ""
				}
				copy(buf[p:], str2[p-i:])
			}
			p = i + m
		}
	}

	dp := make([][]int, n+m)
	for i := range n + m {
		dp[i] = make([]int, m)
		for j := range m {
			dp[i][j] = -1
		}
	}

	// 现在要处理F的情况
	var dfs func(i int, j int) bool

	dfs = func(i int, j int) (ans bool) {
		if i == n+m-1 {
			return true
		}
		if dp[i][j] >= 0 {
			return dp[i][j] > 0
		}
		defer func() {
			if ans {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}()
		// dp[i][j] = -1
		if buf[i] != '-' {
			nj := j
			for nj > 0 && buf[i] != str2[nj] {
				nj = lps[nj-1]
			}
			if buf[i] == str2[nj] {
				nj++
			}
			if nj == m {
				if str1[i-m+1] == 'F' {
					return false
				}
				nj = lps[nj-1]
			}
			ans = dfs(i+1, nj)
		} else {
			for x := range 26 {
				nj := j
				for nj > 0 && int(str2[nj]-'a') != x {
					nj = lps[nj-1]
				}
				if int(str2[nj]-'a') == x {
					nj++
				}
				if nj == m {
					if str1[i-m+1] == 'F' {
						continue
					}
					nj = lps[nj-1]
				}
				ans = dfs(i+1, nj)
				if ans {
					buf[i] = byte(x + 'a')
					break
				}
			}
		}

		return
	}

	if !dfs(0, 0) {
		return ""
	}

	return string(buf)
}

func kmp(p string) []int {
	n := len(p)
	lps := make([]int, n)
	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && p[i] != p[j] {
			j = lps[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		lps[i] = j
	}
	return lps
}
