package p1561

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)
	var i int
	var j = n - 2
	var res int
	for i < j {
		res += piles[j]
		i++
		j -= 2
	}

	return res
}
