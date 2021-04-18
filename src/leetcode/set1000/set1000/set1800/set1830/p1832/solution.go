package p1832

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	var cnt int

	for i := 0; i < len(costs) && costs[i] <= coins; i++ {
		cnt++
		coins -= costs[i]
	}

	return cnt
}
