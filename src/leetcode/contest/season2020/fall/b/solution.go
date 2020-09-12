package b

import "sort"

const MOD = 1000000007

func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	var res int
	var j int
	for i := len(staple) - 1; i >= 0; i-- {
		for j < len(drinks) && staple[i]+drinks[j] <= x {
			j++
		}
		res += j
		if res >= MOD {
			res -= MOD
		}
	}

	return res
}
