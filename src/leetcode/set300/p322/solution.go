package main

import (
	"sort"
	"fmt"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	dp := make([]int, amount + 1)

	for x := 1; x <= amount; x++ {
		dp[x] = amount + 1
		for i := 0; i < len(coins) && coins[i] <= x; i++ {
			if dp[x - coins[i]]+1 < dp[x] {
				dp[x] = dp[x - coins[i]] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
