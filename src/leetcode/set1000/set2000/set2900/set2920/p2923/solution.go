package p2923

import (
	"math"
	"sort"
)

func maxBalancedSubsequenceSum(nums []int) int64 {
	n := len(nums)

	type Pair struct {
		first  int
		second int
	}

	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{nums[i] - i, i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
	})

	arr := make([]int64, 2*n)

	for i := 0; i < 2*n; i++ {
		arr[i] = 0
	}

	set := func(p int, v int64) {
		p += n
		arr[p] = max(arr[p], v)

		for p > 1 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int64 {
		l += n
		r += n
		var res int64
		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}

		return res
	}
	var best int64 = math.MinInt64
	for _, p := range ps {
		tmp := get(0, p.second)
		tmp += int64(p.first + p.second)
		set(p.second, tmp)
		best = max(best, tmp)
	}

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
