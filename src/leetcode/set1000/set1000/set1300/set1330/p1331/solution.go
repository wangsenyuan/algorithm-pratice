package p1331

import "sort"

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{arr[i], i}
	}

	sort.Sort(Pairs(ps))

	res := make([]int, n)

	for i, r := 1, 1; i <= n; i++ {
		res[ps[i-1].second] = r
		if i < n && ps[i].first > ps[i-1].first {
			r++
		}
	}
	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
