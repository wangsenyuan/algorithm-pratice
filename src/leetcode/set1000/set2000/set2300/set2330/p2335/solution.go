package p2335

import "sort"

func fillCups(amount []int) int {
	sort.Ints(amount)

	if amount[0]+amount[1] <= amount[2] {
		return amount[2]
	}

	t := amount[0] + amount[1] - amount[2]
	return (t+1)/2 + amount[2]
}
