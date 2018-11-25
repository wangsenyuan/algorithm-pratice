package p948

import "sort"

func bagOfTokensScore(tokens []int, P int) int {
	n := len(tokens)
	if n == 0 || P == 0 {
		return 0
	}

	sort.Ints(tokens)

	i, j := 0, n

	var ans int
	var points int
	for i < j {
		if P >= tokens[i] {
			points++
			ans = max(ans, points)
			P -= tokens[i]
			i++
		} else if points > 0 {
			points--
			P += tokens[j-1]
			j--
		} else {
			break
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func bagOfTokensScore1(tokens []int, P int) int {
	n := len(tokens)

	if P == 0 || n == 0 {
		return 0
	}

	sort.Ints(tokens)

	if P < tokens[0] {
		return 0
	}

	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(tokens[i])
	}

	ss := make([]int64, n+1)

	for i, j := n-1, 1; i >= 0; i, j = i-1, j+1 {
		ss[j] = ss[j-1] + int64(tokens[i])
	}

	var best int

	for i := 0; i < n; i++ {
		// try to face down first (i + 1) tokens
		a := sum[i+1]
		// need P + a
		j := sort.Search(len(ss), func(j int) bool {
			return ss[j]+int64(P) >= a
		})
		if i+1+j <= n {
			if i+1-j > best {
				best = i + 1 - j
			}
		}
	}

	return best
}
