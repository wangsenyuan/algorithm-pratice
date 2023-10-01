package p121

import (
	"math"
)

func maxProfit(prices []int) int {
	buy, sell := math.MinInt32, 0

	for _, price := range prices {
		if buy < -price {
			buy = -price
		}

		if buy+price > sell {
			sell = buy + price
		}
	}

	return sell
}

func maxProfit1(prices []int) int {
	st := make([]int, len(prices))
	i := 0
	profit := 0

	for _, price := range prices {
		for i > 0 && price < st[i-1] {
			profit = max(st[i-1]-st[0], profit)
			i--
		}
		st[i] = price
		i++
	}

	if i > 0 {
		profit = max(st[i-1]-st[0], profit)
	}

	return profit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
