package main

func maxProfit(prices []int) int {
	profit := 0
	i := 0
	for i < len(prices) {
		price := prices[i]
		j := i
		for j < len(prices)-1 && prices[j+1] >= prices[j] {
			j++
		}
		if j > i {
			profit += prices[j] - price
			i = j + 1
		} else {
			i++
		}
	}

	return profit
}
