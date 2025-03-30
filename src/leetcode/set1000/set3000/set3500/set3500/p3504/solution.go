package p3504

func longestPalindrome(s string, t string) int {
	// 假设 s[:i] 和后缀和t[j:]的前缀，行成了一个回文
	// 且回文的中心在s中， AB是s[:i]的后缀
	// 那么rA 肯定是t[:j]的前缀，B肯定是回文
	// 那么如果在s中找到了A，且找到A后面最长的回文B，就能找到一个答案
	n := len(s)
	m := len(t)

	rs := reverse(s)

	dp := make([][]bool, max(n, m))
	for i := 0; i < max(n, m); i++ {
		dp[i] = make([]bool, max(n, m))
	}
	fp := make([]int, max(n, m))
	for j := 0; j < n; j++ {
		fp[j] = 1
		dp[j][j] = true
		for i := j - 1; i >= 0; i-- {
			if rs[i] == rs[j] && (i+1 == j || dp[i+1][j-1]) {
				dp[i][j] = true
				fp[j] = max(fp[j], j-i+1)
			}
		}
	}

	// 中点在s中的情况
	var best int
	for j := 0; j < m; j++ {
		x := t[j:]
		p := kmp(x)
		var k int
		for i := 0; i < n; i++ {
			for k > 0 && x[k] != rs[i] {
				k = p[k-1]
			}
			if x[k] == rs[i] {
				k++
			}
			// 2 * k + fp[i - k]
			tmp := 2 * k
			if i-k >= 0 {
				tmp += fp[i-k]
			}
			best = max(best, tmp)
			if k == len(x) {
				k = p[k-1]
			}
		}
	}

	// 现在处理中点在t中的情况
	for i := range dp {
		clear(dp[i])
	}
	clear(fp)

	for i := m - 1; i >= 0; i-- {
		fp[i] = 1
		dp[i][i] = true
		for j := i + 1; j < m; j++ {
			if t[i] == t[j] && (i+1 == j || dp[i+1][j-1]) {
				dp[i][j] = true
				fp[j] = max(fp[j], j-i+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		x := reverse(s[:i+1])
		p := kmp(x)
		var k int
		for j := 0; j < m; j++ {
			for k > 0 && x[k] != t[j] {
				k = p[k-1]
			}
			if x[k] == t[j] {
				k++
			}
			tmp := 2 * k
			if j-k >= 0 {
				tmp += fp[j-k]
			}
			best = max(best, tmp)
			if k == len(x) {
				k = p[k-1]
			}
		}
	}

	return best
}

func kmp(x string) []int {
	p := make([]int, len(x))
	for i := 1; i < len(x); i++ {
		j := p[i-1]
		for j > 0 && x[j] != x[i] {
			j = p[j-1]
		}
		if x[j] == x[i] {
			j++
		}
		p[i] = j
	}
	return p
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
