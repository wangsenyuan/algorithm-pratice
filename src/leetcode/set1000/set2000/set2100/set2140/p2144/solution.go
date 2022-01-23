package p2144

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)
	n := len(cost)
	var sum int

	for i := n - 1; i >= 0; {
		sum += cost[i]
		i--
		if i >= 0 {
			sum += cost[i]
			i--
		}
		if i >= 0 {
			i--
		}
	}

	return sum
}
