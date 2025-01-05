package p3413

import "slices"

func maximumCoins(coins [][]int, k int) int64 {
	slices.SortFunc(coins, func(a, b []int) int {
		return a[0] - b[0]
	})

	var res int
	n := len(coins)
	var sum int
	for l, r := 0, 0; l < n; l++ {
		for r < n && coins[r][1]-coins[l][0]+1 <= k {
			sum += coins[r][2] * (coins[r][1] - coins[r][0] + 1)
			r++
		}
		res = max(res, sum)
		if r < n {
			i := max(0, coins[l][0]+k-coins[r][0])
			res = max(res, sum+i*coins[r][2])
		}
		sum -= (coins[l][1] - coins[l][0] + 1) * coins[l][2]
	}

	sum = 0

	for l, r := n-1, n-1; r >= 0; r-- {
		for l >= 0 && coins[r][1]-coins[l][0]+1 <= k {
			sum += (coins[l][1] - coins[l][0] + 1) * coins[l][2]
			l--
		}
		res = max(res, sum)
		if l >= 0 {
			i := max(0, coins[l][1]-(coins[r][1]-k))
			res = max(res, sum+i*coins[l][2])
		}
		sum -= (coins[r][1] - coins[r][0] + 1) * coins[r][2]
	}

	return int64(res)
}
