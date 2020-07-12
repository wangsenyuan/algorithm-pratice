package p309

const N_INF = -(1 << 20)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	if len(prices) == 2 {
		return max(0, prices[1]-prices[0])
	}

	var sell1, sell2 int
	var buy int = N_INF

	var profit int

	for i := 0; i < len(prices); i++ {
		tmpSell2, tmpBuy := sell2, buy
		buy = max(buy, sell1-prices[i])
		sell2 = max(sell2, prices[i]+tmpBuy)
		sell1 = tmpSell2
		profit = max(profit, sell2)
	}

	return profit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
