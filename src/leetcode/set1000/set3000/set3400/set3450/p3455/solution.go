package p3455

func shortestMatchingSubstring(s string, p string) int {
	// 将p分城3段
	var arr []string
	for i, j := 0, 0; i <= len(p); i++ {
		if i == len(p) || p[i] == '*' {
			if j < i {
				arr = append(arr, p[j:i])
			}
			j = i + 1
		}
	}
	if len(arr) == 0 {
		return 0
	}
	// m := len(arr)

	n := len(s)
	// dp[i] 是目前为止，匹配s[...i] 的最右边的j
	dp := make([]int, n)

	for i, cur := range arr {
		lps := kmp(cur)
		var it int
		prev := -n
		fp := make([]int, n)
		for j := 0; j < n; j++ {
			fp[j] = -n
			if j >= len(cur) {
				prev = max(prev, dp[j-len(cur)])
			}

			for it > 0 && cur[it] != s[j] {
				it = lps[it-1]
			}
			if cur[it] == s[j] {
				it++
			}

			if it == len(cur) {
				if i == 0 {
					fp[j] = j - it + 1
				} else {
					fp[j] = prev
				}
				it = lps[it-1]
			}
		}
		copy(dp, fp)
	}
	best := n + 1
	for i := 0; i < n; i++ {
		j := dp[i]
		if j >= 0 {
			best = min(best, i-j+1)
		}
	}

	if best > n {
		return -1
	}

	return best
}

func kmp(p string) []int {
	n := len(p)
	l := make([]int, n)
	for i := 1; i < n; i++ {
		j := l[i-1]
		for j > 0 && p[i] != p[j] {
			j = l[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		l[i] = j
	}
	return l
}
