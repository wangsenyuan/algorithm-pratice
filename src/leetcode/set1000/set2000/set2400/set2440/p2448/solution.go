package p2448

import (
	"math"
	"sort"
)

func minCost(nums []int, cost []int) int64 {
	n := len(nums)
	// 跟定某个x时，(nums[i] - x) * cost[i] = nums[i] * cost[i] - x * cost[i]
	// when x < nums[i]
	type Pair struct {
		first  int
		second int
	}
	pairs := make([]Pair, n)

	for i := 0; i < n; i++ {
		pairs[i] = Pair{nums[i], cost[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].first < pairs[j].first
	})

	cost_suf := make([]int64, n+1)

	var suf int64

	for i := n - 1; i >= 0; i-- {
		cost_suf[i] = cost_suf[i+1] + int64(pairs[i].second)
		suf += int64(nums[i]) * int64(cost[i])
	}
	var res int64 = math.MaxInt64
	var cost_pref int64
	var pref int64
	for i := 0; i < n; i++ {
		// if take p.first as x
		cur := pairs[i]
		suf -= int64(cur.first) * int64(cur.second)
		pref += int64(cur.first) * int64(cur.second)
		cost_pref += int64(cur.second)
		tmp := cost_pref*int64(cur.first) - pref
		tmp += suf - cost_suf[i+1]*int64(cur.first)

		if tmp < res {
			res = tmp
		}
	}
	return res
}
