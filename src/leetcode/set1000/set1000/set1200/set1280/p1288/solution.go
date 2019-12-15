package p1288

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	n := len(intervals)

	if n == 0 {
		return 0
	}

	its := make([]Pair, n)

	for i := 0; i < n; i++ {
		its[i] = Pair{intervals[i][0], intervals[i][1]}
	}

	sort.Sort(Pairs(its))

	// left := its[0].first
	right := its[0].second
	cnt := n
	for i := 1; i < n; i++ {
		if its[i].second <= right {
			cnt--
			continue
		}
		// left = its[i].first
		right = its[i].second
	}

	return cnt
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || (ps[i].first == ps[j].second && ps[i].second > ps[j].second)
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
