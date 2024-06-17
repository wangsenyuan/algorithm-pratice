package p3185

import "sort"

type Pair struct {
	first  int
	second int
}

func maximumTotalDamage(power []int) int64 {
	sort.Ints(power)

	var ps []Pair

	n := len(power)

	for i := 0; i < n; {
		var sum int
		j := i
		for i < n && power[j] == power[i] {
			sum += power[i]
			i++
		}
		ps = append(ps, Pair{sum, power[j]})
	}

	n = len(ps)

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = ps[i].first
		for j := i - 1; j >= 0; j-- {
			if ps[j].second+2 < ps[i].second {
				dp[i] = max(dp[i], dp[j]+ps[i].first)
				break
			}
		}
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
	}

	return int64(dp[n-1])
}
