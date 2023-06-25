package p2747

const inf = 1 << 30

func minimizeConcatenatedLength(words []string) int {
	// dp[a][b] 表示前以a开始，并以b结尾的字符串的最小长度
	dp := make([]int, 26*26)
	fp := make([]int, 26*26)
	for i := 0; i < len(dp); i++ {
		dp[i] = inf
		fp[i] = inf
	}
	dp[first(words[0])*26+last(words[0])] = len(words[0])

	n := len(words)

	for i := 1; i < n; i++ {
		cur := words[i]
		x := first(cur)
		y := last(cur)
		ln := len(cur)

		for a := 0; a < 26; a++ {
			for b := 0; b < 26; b++ {
				tmp1 := ln + dp[a*26+b]
				tmp2 := tmp1
				if y == a {
					tmp1--
				}
				if x == b {
					tmp2--
				}
				fp[x*26+b] = min(fp[x*26+b], tmp1)
				fp[a*26+y] = min(fp[a*26+y], tmp2)
			}
		}
		for j := 0; j < len(fp); j++ {
			dp[j] = fp[j]
			fp[j] = inf
		}
	}
	best := inf
	for i := 0; i < len(dp); i++ {
		best = min(best, dp[i])
	}
	return best
}

func first(a string) int {
	return int(a[0] - 'a')
}

func last(a string) int {
	return int(a[len(a)-1] - 'a')
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
