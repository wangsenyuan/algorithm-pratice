package main

import "fmt"

func maxProfit1(prices []int, fee int) int {
	n := len(prices)

	dp := make([]int, n)
	for i := 1; i < n; i++ {
		// sell at day i
		profit := 0
		for j := 0; j < i; j++ {
			// buy at day j
			tmp := prices[i] - prices[j] - fee
			if j > 0 {
				tmp += dp[j-1]
			}
			if profit < tmp {
				profit = tmp
			}
		}
		dp[i] = max(profit, dp[i-1])
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	buy := make([]int, n)
	sell := make([]int, n)
	skip := make([]int, n)
	hold := make([]int, n)

	buy[0] = -prices[0]
	hold[0] = - prices[0]

	for i := 1; i < n; i++ {
		buy[i] = max(skip[i-1], sell[i-1]) - prices[i]
		hold[i] = max(buy[i-1], hold[i-1])
		skip[i] = max(sell[i-1], skip[i-1])
		sell[i] = max(buy[i-1], hold[i-1]) + prices[i] - fee
	}

	ans := 0
	ans = max(ans, buy[n-1])
	ans = max(ans, sell[n-1])
	ans = max(ans, skip[n-1])
	ans = max(ans, hold[n-1])
	return ans
}

func main() {
	//prices := []int{1, 3, 2, 8, 4, 9}
	prices := []int{1, 3, 7, 5, 10, 3}
	fmt.Println(maxProfit1(prices, 3))
}
