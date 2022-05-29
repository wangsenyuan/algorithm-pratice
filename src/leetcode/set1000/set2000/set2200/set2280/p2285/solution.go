package p2285

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	deg := make([]int, n)
	for _, road := range roads {
		a, b := road[0], road[1]
		deg[a]++
		deg[b]++
	}
	sort.Ints(deg)
	var res int64

	for i := n; i > 0; i-- {
		res += int64(i) * int64(deg[i-1])
	}
	return res
}
