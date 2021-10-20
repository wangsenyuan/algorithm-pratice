package lcp49

import "sort"

func ringGame(challenge []int64) int64 {
	n := len(challenge)

	level := make([]int, n)

	for i := 0; i < n; i++ {
		level[i] = i
	}

	sort.Slice(level, func(i, j int) bool {
		return challenge[level[i]] < challenge[level[j]]
	})

	check := func(x int, score int64) bool {
		score |= challenge[x]
		l, r := x, x
		for {
			ll := (l - 1 + n) % n
			rr := (r + 1) % n
			if l == rr || r == ll {
				return true
			}
			if score >= challenge[ll] {
				score |= challenge[ll]
				l = ll
			} else if score >= challenge[rr] {
				score |= challenge[rr]
				r = rr
			} else {
				return false
			}
		}
	}

	maxScore := challenge[level[n-1]]
	var start int64 = 1
	tmp := maxScore
	for tmp > 0 {
		start <<= 1
		tmp >>= 1
	}
	start >>= 1

	for _, l := range level {
		minScore := max(start, challenge[l])
		for minScore < maxScore {
			score := (minScore + maxScore) / 2

			if check(l, score) {
				maxScore = score
			} else {
				minScore = score + 1
			}
		}
	}
	return maxScore
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
