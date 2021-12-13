package p2075

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	qs := make([]Query, len(queries))

	for i, cur := range queries {
		qs[i] = Query{cur, i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].query < qs[j].query
	})

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	ans := make([]int, len(qs))
	var best int
	for i, j := 0, 0; i < len(qs); i++ {
		cur := qs[i]
		for j < len(items) && items[j][0] <= cur.query {
			best = max(best, items[j][1])
			j++
		}
		ans[cur.index] = best
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Query struct {
	query int
	index int
}
