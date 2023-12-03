package p2953

import "sort"

func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	n := len(coins)
	var sum int
	var res int
	for i := 0; i < n; i++ {
		for coins[i] > sum+1 {
			res++
			sum += sum + 1
		}
		sum += coins[i]
	}

	for target > sum {
		res++
		sum += sum + 1
	}

	return res
}
