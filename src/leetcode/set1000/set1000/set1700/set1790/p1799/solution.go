package p1799

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	var prev int
	for i := 0; i < len(coins); i++ {
		if prev+1 < coins[i] {
			return prev + 1
		}
		prev += coins[i]
	}
	return prev + 1
}
