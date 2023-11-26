package p2947

import (
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	p := make([]Pair, n)
	for i := 0; i < n; i++ {
		p[i] = Pair{nums[i], i}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].first < p[j].first || p[i].first == p[j].first && p[i].second < p[j].second
	})

	res := make([]int, n)
	buf := make([]int, n)
	for i := 0; i < n; {
		j := i
		for i < n {
			buf[i] = p[i].second
			i++
			if i == n || p[i].first-p[i-1].first > limit {
				break
			}
		}
		sort.Ints(buf[j:i])
		for j < i {
			res[buf[j]] = p[j].first
			j++
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}
