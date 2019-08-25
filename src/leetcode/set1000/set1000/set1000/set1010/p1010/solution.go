package p1010

import "sort"

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)

	sort.Sort(Pairs(costs))

	// city B costs less that city A go left
	var res int
	for i := 0; i < n; i++ {
		if 2*i < n {
			// first half
			res += costs[i][1]
		} else {
			res += costs[i][0]
		}
	}
	return res
}

type Pairs [][]int

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	a := ps[i]
	b := ps[j]
	return a[1]-a[0] < b[1]-b[0]
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
